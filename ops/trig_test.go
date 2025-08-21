package ops_test

import (
	"math"
	"testing"

	"github.com/kelledge/uom/ops"
	"github.com/kelledge/uom/std"
	"github.com/stretchr/testify/require"
)

func TestTrigBasic(t *testing.T) {
	const eps = 1e-12

	t.Run("unit circle values", func(t *testing.T) {
		type angleCase struct {
			label string
			theta float64
			sin   float64
			cos   float64
			tan   float64 // TODO: is less straight forward that one would think
		}

		cases := []angleCase{
			{"0", 0, 0, 1, 0},
			{"π/4", math.Pi / 4, math.Sqrt(2) / 2, math.Sqrt(2) / 2, 1},
			{"π/2", math.Pi / 2, 1, 0, math.Inf(1)},
			{"π", math.Pi, 0, -1, 0},
			{"3π/2", 3 * math.Pi / 2, -1, 0, math.Inf(-1)},
		}

		for _, c := range cases {
			theta := std.Radian.Of(c.theta)

			require.InDeltaf(t, c.sin, ops.Sin(theta).Value(), eps, "sin(%s) failed", c.label)
			require.InDeltaf(t, c.cos, ops.Cos(theta).Value(), eps, "cos(%s) failed", c.label)
		}
	})
}
