package usc_test

import (
	"testing"

	"github.com/kelledge/uom/dim"
	"github.com/kelledge/uom/si"
	"github.com/kelledge/uom/usc"
	"github.com/stretchr/testify/require"
)

func TestSpotCheck(t *testing.T) {
	const delta = 1e-15

	t.Run("gallon", func(t *testing.T) {
		require.Equal(t, dim.Volume{}.Dimension(), usc.Gallon.Dimension())
		require.InDelta(t, 0.003785411784, usc.Gallon.ConversionFactor(), delta)
	})

	t.Run("standard gravity is 9.81 m/s^2", func(t *testing.T) {
		gUSC := usc.StandardGravity.Of(1).Canonical()
		gSI := si.StandardGravity.Of(1).Canonical()

		require.InDelta(t, gSI, gUSC, 1e-12)
	})
}
