package find_minimum_in_rotated_sorted_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findMin(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		require.Equal(t, 1, findMin([]int{3, 4, 5, 1, 2}))
	})

	t.Run("case 2", func(t *testing.T) {
		require.Equal(t, 0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	})

	t.Run("case 3", func(t *testing.T) {
		require.Equal(t, 11, findMin([]int{11, 13, 15, 17}))
	})
}
