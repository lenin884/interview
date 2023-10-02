package evaluate_division

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		equations := [][]string{{"a", "b"}, {"b", "c"}}
		values := []float64{2.0, 3.0}
		queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
		expected := []float64{6.0, 0.5, -1.0, 1.0, -1.0}
		actual := calcEquation(equations, values, queries)
		require.Equal(t, expected, actual)
	})
}
