package rotate_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("1 element", func(t *testing.T) {
		nums := []int{1}
		k := 1

		rotate(nums, k)
		want := []int{1}

		require.Equal(t, want, nums)
	})

	t.Run("2 elements", func(t *testing.T) {
		nums := []int{1, 2}
		k := 1

		rotate(nums, k)
		want := []int{2, 1}

		require.Equal(t, want, nums)
	})

	t.Run("10 elements", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
		k := 3

		rotate(nums, k)
		want := []int{6, 7, 8, 1, 2, 3, 4, 5}

		require.Equal(t, want, nums)
	})

	t.Run("4 elements k = 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		k := 2

		rotate(nums, k)
		want := []int{3, 4, 1, 2}

		require.Equal(t, want, nums)
	})
}
