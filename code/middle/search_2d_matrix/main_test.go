package search_2d_matrix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_search2DMatrix(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
		target := 3
		require.Equal(t, true, searchMatrix(matrix, target))
	})

	t.Run("case 2", func(t *testing.T) {
		matrix := [][]int{{1}, {2}, {5}}
		target := 5
		require.Equal(t, true, searchMatrix(matrix, target))
	})

	t.Run("case 3", func(t *testing.T) {
		matrix := [][]int{{1}, {3}}
		target := 2
		require.Equal(t, false, searchMatrix(matrix, target))
	})
}
