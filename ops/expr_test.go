package ops_test

import (
	"testing"

	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
	"github.com/kelledge/uom/ops"
	"github.com/kelledge/uom/si"
	"github.com/kelledge/uom/std"
	"github.com/kelledge/uom/usc"
	"github.com/stretchr/testify/require"
)

func TestExpr(t *testing.T) {
	oneKg := si.Kilogram.Of(1)
	oneM := si.Meter.Of(1)
	twoM := si.Meter.Of(2)
	oneS := std.Second.Of(1)

	oneMpS2 := si.MeterPerSecondSquared.Of(1)

	t.Run("add enforces dimension rules", func(t *testing.T) {
		require.NotPanics(t, func() {
			actual := ops.E[dim.Length](oneM).Add(twoM).As(si.Meter).Value()
			require.Equal(t, 3.0, actual)
		})
		require.Panics(t, func() {
			ops.E[dim.Length](oneM).Add(oneS)
		})
	})

	t.Run("sub enforces dimension rules", func(t *testing.T) {
		require.NotPanics(t, func() {
			actual := ops.E[dim.Length](twoM).Sub(oneM).As(si.Meter).Value()
			require.Equal(t, 1.0, actual)
		})
		require.Panics(t, func() {
			ops.E[dim.Length](oneM).Sub(oneS)
		})
	})

	t.Run("mul combines dimension properly", func(t *testing.T) {
		require.NotPanics(t, func() {
			actual := ops.E[dim.Force](oneKg).Mul(oneMpS2).As(si.Newton).Value()
			require.Equal(t, 1.0, actual)
		})
	})

	t.Run("div combines dimension properly", func(t *testing.T) {
		require.NotPanics(t, func() {
			actual := ops.E[dim.Velocity](oneM).Div(oneS).As(si.MeterPerSecond).Value()
			require.Equal(t, 1.0, actual)
		})
	})

	t.Run("pow modifies dimension properly", func(t *testing.T) {
		require.NotPanics(t, func() {
			actual := ops.E[dim.Area](oneM).Pow(2).As(si.MeterSquared).Value()
			require.Equal(t, 1.0, actual)
		})
	})

	t.Run("inv modifies dimension properly", func(t *testing.T) {
		require.NotPanics(t, func() {
			actual := ops.E[dim.Frequency](oneS).Inv().As(std.Herz).Value()
			require.Equal(t, 1.0, actual)
		})
	})

	t.Run("pow handles raising to the zero power", func(t *testing.T) {
		require.NotPanics(t, func() {
			actual := ops.E[dim.Dimensionless](twoM).Pow(0).As(std.Scalar).Value()
			require.Equal(t, 1.0, actual)
		})
	})

	t.Run("pow panics when raising zero to the zero power", func(t *testing.T) {
		require.Panics(t, func() {
			ops.E[dim.Dimensionless](std.Scalar.Of(0)).Pow(0)
		})
	})

	t.Run("inv panics when value is zero", func(t *testing.T) {
		require.Panics(t, func() {
			ops.E[dim.Dimensionless](std.Scalar.Of(0)).Inv()
		})
	})

	t.Run("add requires operands", func(t *testing.T) {
		require.Panics(t, func() { ops.E[dim.Length](oneM).Add(nil) })
	})

	t.Run("sub requires operands", func(t *testing.T) {
		require.Panics(t, func() { ops.E[dim.Length](oneM).Sub(nil) })
	})

	t.Run("mul requires operands", func(t *testing.T) {
		require.Panics(t, func() { ops.E[dim.Length](oneM).Mul(nil) })
	})

	t.Run("div requires operands", func(t *testing.T) {
		require.Panics(t, func() { ops.E[dim.Length](oneM).Div(nil) })
	})

	t.Run("add enforces no nil operands", func(t *testing.T) {
		require.Panics(t, func() { ops.E[dim.Length](oneM).Add() })
	})

	t.Run("sub enforces no nil operands", func(t *testing.T) {
		require.Panics(t, func() { ops.E[dim.Length](oneM).Sub() })
	})

	t.Run("mul enforces no nil operands", func(t *testing.T) {
		require.Panics(t, func() { ops.E[dim.Length](oneM).Mul() })
	})

	t.Run("div enforces no nil operands", func(t *testing.T) {
		require.Panics(t, func() { ops.E[dim.Length](oneM).Div() })
	})

	t.Run("operates on quantites and expressions", func(t *testing.T) {
		// 2 m/s
		speed := ops.E[dim.Velocity](twoM).Div(oneS)

		// 1 m / 2 m/s => 0.5 s
		actual := ops.E[dim.Time](oneM).Div(speed).As(std.Second).Value()

		require.Equal(t, 0.5, actual)
	})

	t.Run("dimension times its inverse is dimensionless", func(t *testing.T) {
		actual := ops.E[dim.Dimensionless](oneM).Inv().Mul(oneM).As(std.Scalar).Value()
		require.Equal(t, 1.0, actual)
	})

	t.Run("can call canonical to not hardcode units", func(t *testing.T) {
		uom.DefaultRegistry.UseCanonicalSet(si.MKS)

		in := usc.Inch.Of(1)
		ft := usc.Foot.Of(1)

		a := ops.E[dim.Length](in).Mul(std.Scalar.Of(36)).AsCanonical().Value()
		b := ops.E[dim.Length](ft).Mul(std.Scalar.Of(3)).AsCanonical().Value()

		require.InDelta(t, a, b, 1e-12)
	})
}
