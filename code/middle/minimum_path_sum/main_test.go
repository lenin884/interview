package minimum_path_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_minimumPathSum(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		grid := [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}
		want := 7
		got := minPathSum(grid)
		require.Equal(t, want, got)
	})
}
