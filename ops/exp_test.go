package ops_test

import (
	"math"
	"testing"

	"github.com/kelledge/uom/ops"
	"github.com/kelledge/uom/std"
	"github.com/stretchr/testify/require"
)

func TestExpLogBasic(t *testing.T) {
	const eps = 1e-12

	t.Run("exp and ln are inverses (where defined)", func(t *testing.T) {
		expected := std.Scalar.Of(math.Pi)
		actual := ops.Ln(ops.Exp(expected))

		require.InDelta(t, expected.Value(), actual.Value(), eps)
	})

	t.Run("basic identities", func(t *testing.T) {
		require.InDelta(t, 1.0, ops.Exp(std.Scalar.Of(0)).Value(), eps)
		require.InDelta(t, 0.0, ops.Ln(std.Scalar.Of(1)).Value(), eps)
		require.InDelta(t, 0.0, ops.Log10(std.Scalar.Of(1)).Value(), eps)

		b := std.Scalar.Of(2.5)
		require.InDelta(t, 1.0, ops.Log(b, b).Value(), eps)

		require.InDelta(t, 1.0, ops.Log10(std.Scalar.Of(10)).Value(), eps)
	})
}
