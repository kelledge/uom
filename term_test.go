package uom_test

import (
	"testing"

	"github.com/kelledge/uom"
	"github.com/stretchr/testify/require"
)

func TestUnitTermOperations(t *testing.T) {
	u := uom.DefineUnit[uom.Length]("inch", 0.0254)

	t.Run("test per", func(t *testing.T) {
		inchPer := uom.U(u).Per()

		actual := uom.DefaultDimCodec.Decode(inchPer.Dimension())

		require.Equal(t, -1, actual.Length)
		require.Equal(t, 1/0.0254, inchPer.ConversionFactor())
	})

	t.Run("test pow", func(t *testing.T) {
		inchCubed := uom.U(u).Pow(3)

		actual := uom.DefaultDimCodec.Decode(inchCubed.Dimension())

		require.Equal(t, 3, actual.Length)
		require.Equal(t, 0.0254*0.0254*0.0254, inchCubed.ConversionFactor())
	})
}
