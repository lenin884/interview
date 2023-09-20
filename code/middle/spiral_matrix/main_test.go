package spiral_matrix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12}}
		want := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
		require.Equal(t, want, spiralOrder(matrix))
	})

	t.Run("test 2", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9}}
		want := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
		require.Equal(t, want, spiralOrder(matrix))
	})
}
