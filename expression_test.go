package uom_test

import (
	"testing"

	"github.com/kelledge/uom"
	"github.com/stretchr/testify/require"
)

func TestExpr(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		kilogram := uom.DefineUnit[uom.Mass]("kilogram", 1.0)
		meter := uom.DefineUnit[uom.Length]("meter", 1.0)
		second := uom.DefineUnit[uom.Time]("second", 1.0)

		newton := uom.DefineUnit[uom.Force]("newton", 1.0)
		lbf := uom.DefineUnit[uom.Force]("lbf", 4.4482216152605)

		q1 := uom.Q(1, kilogram)
		q2 := uom.Q(1, meter)
		q3 := uom.Q(1, second)

		// force := func(m uom.Quantity[uom.Mass], a uom.Quantity[uom.Acceleration]) uom.Quantity[uom.Force] {

		// }

		// 1 kg * 1 m / s^2
		e := uom.One().Mul(q1, q2).Div(q3, q3)

		uom.As(e, meter)

		require.Equal(t, 1.0, uom.As(e, newton).Value())
		require.Equal(t, 1/4.4482216152605, uom.As(e, lbf).Value())
	})
}
