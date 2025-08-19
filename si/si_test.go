package si_test

import (
	"testing"

	"github.com/kelledge/uom/dim"
	"github.com/kelledge/uom/si"
	"github.com/stretchr/testify/require"
)

func TestSpotCheck(t *testing.T) {
	t.Run("Kilometer", func(t *testing.T) {
		require.Equal(t, dim.Length{}.Dimension(), si.Kilometer.Dimension(), "Kilometer should have Length dimension")
		require.Equal(t, 1000.0, si.Kilometer.ConversionFactor(), "Kilometer should be 1000 m")
	})

	t.Run("Centimeter", func(t *testing.T) {
		require.Equal(t, dim.Length{}.Dimension(), si.Centimeter.Dimension(), "Centimeter should have Length dimension")
		require.Equal(t, 0.01, si.Centimeter.ConversionFactor(), "Centimeter should be 0.01 m")
	})

	t.Run("Gram", func(t *testing.T) {
		require.Equal(t, dim.Mass{}.Dimension(), si.Gram.Dimension(), "Gram should have Mass dimension")
		require.Equal(t, 0.001, si.Gram.ConversionFactor(), "Gram should be 1e-3 kg")
	})
}
