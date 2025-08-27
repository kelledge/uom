package uom_test

import (
	"encoding/json"
	"testing"

	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
	"github.com/kelledge/uom/si"
	"github.com/kelledge/uom/std"
	"github.com/stretchr/testify/require"
)

type TestProperties struct {
	Length   uom.Quantity[dim.Length] `json:"length"`
	Duration uom.Quantity[dim.Time]   `json:"duration"`
}

func TestQuantity(t *testing.T) {
	t.Run("serialization to json is reversible", func(t *testing.T) {
		expected := TestProperties{
			Length:   si.Meter.Of(10),
			Duration: std.Minute.Of(45),
		}

		input, err := json.Marshal(&expected)
		require.NoError(t, err)

		var actual TestProperties
		err = json.Unmarshal(input, &actual)
		require.NoError(t, err)

		require.Equal(t, expected, actual)
	})
}
