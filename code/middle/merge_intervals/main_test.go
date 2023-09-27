package merge_intervals

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		got := merge(intervals)
		want := [][]int{{1, 6}, {8, 10}, {15, 18}}
		require.Equal(t, want, got)
	})

	t.Run("test 2", func(t *testing.T) {
		intervals := [][]int{{1, 4}, {4, 5}}
		got := merge(intervals)
		want := [][]int{{1, 5}}
		require.Equal(t, want, got)
	})

	t.Run("test 3 largest number", func(t *testing.T) {
		intervals := [][]int{{1, 4}, {0, 4}}
		got := merge(intervals)
		want := [][]int{{0, 4}}
		require.Equal(t, want, got)
	})
}

func TestQuickSort(t *testing.T) {
	t.Run("test quick sort", func(t *testing.T) {
		nums := []int{3, 2, 1, 5, 6, 4}
		quickSort(nums, 0, len(nums)-1)
		want := []int{1, 2, 3, 4, 5, 6}
		require.Equal(t, want, nums)
	})

	t.Run("test quick sort 2", func(t *testing.T) {
		nums := []int{2, 5, 2, 3, 3}
		quickSort(nums, 0, len(nums)-1)
		want := []int{2, 2, 3, 3, 5}
		require.Equal(t, want, nums)
	})
}
