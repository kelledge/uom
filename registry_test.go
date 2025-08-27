package uom_test

import (
	"testing"

	"github.com/kelledge/uom"
	"github.com/stretchr/testify/require"
)

func TestRegistry(t *testing.T) {
	t.Run("duplicate alias for dimension panics", func(t *testing.T) {
		r := uom.NewRegistry()
		var dim uom.DimInt

		require.NotPanics(t, func() {
			r.RegisterUnit(dim, uom.UnitTerm{}, "m")
		})

		require.Panics(t, func() {
			r.RegisterUnit(dim, uom.UnitTerm{}, "m")
		})
	})

	t.Run("empty alias panics", func(t *testing.T) {
		r := uom.NewRegistry()
		var dim uom.DimInt

		require.Panics(t, func() {
			r.RegisterUnit(dim, uom.UnitTerm{}, "")
		})
	})
}
