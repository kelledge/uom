package uom_test

import (
	"testing"

	"github.com/kelledge/uom"
	"github.com/stretchr/testify/require"
)

func TestDimInt(t *testing.T) {
	codec := uom.MustDimCodec(uom.DimSizes{20, 20, 20, 20, 20, 20, 20, 20})

	// Small, always-valid test vectors
	u := uom.DimensionSpec{
		Length: 1,
		Time:   -1,
	}
	v := uom.DimensionSpec{
		Length:      -1,
		Temperature: 1,
	}

	// Encode
	du := codec.MustEncode(u)
	dv := codec.MustEncode(v)

	t.Run("is reversible", func(t *testing.T) {
		expected := uom.DimensionSpec{
			Mass:   1,
			Length: 1,
			Time:   -2,
		}

		dimint := codec.MustEncode(expected)

		actual := codec.Decode(dimint)

		require.Equal(t, expected, actual)
	})

	t.Run("dimint( u + v ) = dimint(u) + dimint(v)", func(t *testing.T) {
		sumSpec := uom.DimensionSpec{
			Length:      u.Length + v.Length,
			Time:        u.Time + v.Time,
			Temperature: u.Temperature + v.Temperature,
			Mass:        u.Mass + v.Mass,
			Current:     u.Current + v.Current,
			Substance:   u.Substance + v.Substance,
			Luminosity:  u.Luminosity + v.Luminosity,
			Money:       u.Money + v.Money,
		}
		require.Equal(t, codec.MustEncode(sumSpec), du+dv)
	})

	t.Run("dimint( u - v ) = dimint(u) - dimint(v)", func(t *testing.T) {
		diffSpec := uom.DimensionSpec{
			Length:      u.Length - v.Length,
			Time:        u.Time - v.Time,
			Temperature: u.Temperature - v.Temperature,
			Mass:        u.Mass - v.Mass,
			Current:     u.Current - v.Current,
			Substance:   u.Substance - v.Substance,
			Luminosity:  u.Luminosity - v.Luminosity,
			Money:       u.Money - v.Money,
		}
		require.Equal(t, codec.MustEncode(diffSpec), du-dv)
	})

	t.Run("dimint( - u ) = - dimint(u)", func(t *testing.T) {
		negSpec := uom.DimensionSpec{
			Length:      -u.Length,
			Time:        -u.Time,
			Temperature: -u.Temperature,
			Mass:        -u.Mass,
			Current:     -u.Current,
			Substance:   -u.Substance,
			Luminosity:  -u.Luminosity,
			Money:       -u.Money,
		}
		require.Equal(t, codec.MustEncode(negSpec), -du)
	})
}

func TestDimSize(t *testing.T) {
	codec := uom.MustDimCodec(uom.DimSizes{20, 20, 20, 10, 10, 10, 10, 10})

	t.Run("dimvals matches paper", func(t *testing.T) {
		expected := [8]int64{1, 20, 400, 8000, 80000, 800000, 8000000, 80000000}
		require.Equal(t, expected, codec.Vals)
	})

	t.Run("dimbias matches paper", func(t *testing.T) {
		expected := int64(444444210)
		require.Equal(t, expected, codec.Bias)
	})
}

func TestDimCodec(t *testing.T) {
	t.Run("enforces sizes are greater than 1", func(t *testing.T) {
		_, err := uom.NewDimCodec(uom.DimSizes{1, 4, 4, 4, 4, 4, 4, 4})
		require.Error(t, err)
	})

	t.Run("enforces sizes are even", func(t *testing.T) {
		_, err := uom.NewDimCodec(uom.DimSizes{5, 4, 4, 4, 4, 4, 4, 4})
		require.Error(t, err)
	})

	t.Run("enforces exponent sizes", func(t *testing.T) {
		codec, err := uom.NewDimCodec(uom.DimSizes{4, 4, 4, 4, 4, 4, 4, 4})
		require.NoError(t, err)

		_, err = codec.Encode(uom.DimensionSpec{
			Length: 1,
		})

		require.NoError(t, err)

		_, err = codec.Encode(uom.DimensionSpec{
			Length: 2,
		})

		require.Error(t, err)
	})
}
