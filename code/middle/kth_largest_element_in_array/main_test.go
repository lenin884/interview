package kth_largest_element_in_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_kth(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{3, 2, 1, 5, 6, 4}
		k := 2
		require.Equal(t, 5, findKthLargest(nums, k))
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
		k := 4
		require.Equal(t, 4, findKthLargest(nums, k))
	})
}
