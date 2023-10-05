package maximum_sum_circular_subarray

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxSubarraySumCircular(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{1, -2, 3, -2}
		want := 3
		got := maxSubarraySumCircular(nums)
		require.Equal(t, want, got)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{5, -3, 5}
		want := 10
		got := maxSubarraySumCircular(nums)
		require.Equal(t, want, got)
	})

	t.Run("case 3", func(t *testing.T) {
		nums := []int{-3, -1, -2, -1}
		want := -1
		got := maxSubarraySumCircular(nums)
		require.Equal(t, want, got)
	})
}
