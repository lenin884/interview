package combinations

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCombine(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		require.Equal(t, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}, combine(4, 2))
	})
}
