package search_in_rotated_sorted_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_search(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 0
		require.Equal(t, 4, search(nums, target))
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 3
		require.Equal(t, -1, search(nums, target))
	})
}
