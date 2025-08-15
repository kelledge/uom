package uom_test

import (
	"testing"

	"github.com/kelledge/uom"
	"github.com/stretchr/testify/require"
)

func TestDeriveUnit(t *testing.T) {
	t.Run("panics when terms do not match dimension type argument", func(t *testing.T) {
		meter := uom.DefineUnit[uom.Length]("meter", 1.0)
		second := uom.DefineUnit[uom.Time]("second", 1.0)

		require.Panics(t, func() {
			_ = uom.DeriveUnit[uom.Acceleration]("meter/second", uom.U(meter), uom.U(second).Per())
		})
	})
}
