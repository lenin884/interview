package find_first_and_last_position_in_sorted_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_searchRanges(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 8
		require.Equal(t, []int{3, 4}, searchRange(nums, target))
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 6
		require.Equal(t, []int{-1, -1}, searchRange(nums, target))
	})
}
