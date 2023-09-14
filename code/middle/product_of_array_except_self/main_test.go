package product_of_array_except_self

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("4 elements", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		expected := []int{24, 12, 8, 6}

		require.Equal(t, expected, productExceptSelf(nums))
	})

	t.Run("5 elements", func(t *testing.T) {
		nums := []int{-1, 1, 0, -3, 3}
		expected := []int{0, 0, 9, 0, 0}

		require.Equal(t, expected, productExceptSelf(nums))
	})
}
