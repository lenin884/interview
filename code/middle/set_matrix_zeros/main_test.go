package set_matrix_zeros

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 0, 6},
			{7, 8, 9}}
		want := [][]int{
			{1, 0, 3},
			{0, 0, 0},
			{7, 0, 9}}
		setZeroesHashMap(matrix)
		require.Equal(t, want, matrix)
	})

	t.Run("test 2", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 0, 6},
			{7, 8, 9}}
		want := [][]int{
			{1, 0, 3},
			{0, 0, 0},
			{7, 0, 9}}
		setZeroes(matrix)
		require.Equal(t, want, matrix)
	})
}
