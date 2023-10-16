package unique_path_2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_uniquePath(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		obstacleGrid := [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0}}

		got := uniquePathsWithObstacles(obstacleGrid)
		want := 2
		require.Equal(t, want, got)
	})
}
