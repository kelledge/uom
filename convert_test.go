package uom_test

import (
	"testing"

	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
	"github.com/kelledge/uom/si"
	"github.com/kelledge/uom/std"
	"github.com/kelledge/uom/usc"
	"github.com/stretchr/testify/require"
)

type TestDim struct{}

func (TestDim) Dimension() uom.DimInt { return 42 }
func (TestDim) Name() string          { return "Test" }

func TestAsUnit(t *testing.T) {
	t.Run("panics on dimension mismatch", func(t *testing.T) {
		term := uom.U(si.Meter)

		require.Panics(t, func() {
			_ = uom.AsUnit[dim.Time](term)
		})
	})
}

func TestAs(t *testing.T) {
	t.Run("panics when there is a dimension mismatch", func(t *testing.T) {
		q := si.Meter.Of(1)

		require.Panics(t, func() {
			_ = uom.As(q, std.Second)
		})
	})
}

func TestAsCanonical(t *testing.T) {
	t.Run("meters is canonical in MKS system", func(t *testing.T) {
		uom.DefaultRegistry.UseCanonicalSet(si.MKS)

		q := si.Centimeter.Of(100)
		canon := uom.AsCanonical[dim.Length](q)

		require.InDelta(t, 1, canon.Value(), 1e-9)
	})

	t.Run("panics when no canonical set", func(t *testing.T) {
		uom.DefaultRegistry.UseCanonicalSet(nil)

		val := usc.Foot.Of(1)
		require.Panics(t, func() {
			_ = uom.AsCanonical[dim.Length](val)
		})
	})

	t.Run("panics when no canonical unit for dimension", func(t *testing.T) {
		uom.DefaultRegistry.UseCanonicalSet(si.MKS)

		u := uom.DefineUnit[TestDim]("nonsense", 1.0)

		require.Panics(t, func() {
			_ = uom.AsCanonical[TestDim](u.Of(5))
		})
	})
}
