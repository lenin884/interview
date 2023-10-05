package permutations

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPermute(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 2, 1},
			{3, 1, 2},
		}

		require.Equal(t, expected, permute(input))
	})
}
