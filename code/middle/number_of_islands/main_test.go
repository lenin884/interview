package number_of_islands

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_numIslands(t *testing.T) {
	t.Run("should return 0 for empty grid", func(t *testing.T) {
		require.Equal(t, 0, numIslands([][]byte{}))
	})

	t.Run("case 1", func(t *testing.T) {
		grid := [][]byte{
			[]byte("11110"),
			[]byte("11010"),
			[]byte("11000"),
			[]byte("00000"),
		}
		require.Equal(t, 1, numIslands(grid))
	})

	t.Run("case 2", func(t *testing.T) {
		grid := [][]byte{
			[]byte("11000"),
			[]byte("11000"),
			[]byte("00100"),
			[]byte("00011"),
		}
		require.Equal(t, 3, numIslands(grid))
	})
}
