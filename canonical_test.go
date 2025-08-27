package uom_test

import (
	"testing"

	"github.com/kelledge/uom"
	"github.com/kelledge/uom/si"
	"github.com/stretchr/testify/require"
)

func TestCanonicalSet(t *testing.T) {
	t.Run("new panics when on dimension collision", func(t *testing.T) {
		require.Panics(t, func() {
			uom.NewCanonicalSet(uom.U(si.Meter), uom.U(si.Meter))
		})
	})

	t.Run("empty canonical set is not valid and panics", func(t *testing.T) {
		require.Panics(t, func() {
			uom.NewCanonicalSet()
		})
	})
}
