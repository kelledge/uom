package uom_test

import (
	"testing"

	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
	"github.com/kelledge/uom/si"
	"github.com/stretchr/testify/require"
)

func TestDeriveUnit(t *testing.T) {
	t.Run("panics when terms do not match dimension type argument", func(t *testing.T) {
		require.Panics(t, func() {
			_ = uom.DeriveUnit[dim.Acceleration]("meter/second", uom.U(si.Meter), uom.U(si.Meter).Per())
		})
	})
}
