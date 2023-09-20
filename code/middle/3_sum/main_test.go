package __sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		input := []int{-1, 0, 1, 2, -1, -4}
		expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}

		require.Equal(t, expected, threeSum(input))
	})

	t.Run("test 2", func(t *testing.T) {
		input := []int{}
		expected := [][]int{}

		require.Equal(t, expected, threeSum(input))
	})

	t.Run("test 3 large zero nums", func(t *testing.T) {
		input := []int{0, 0, 0, -3, 3, 0}
		expected := [][]int{{-3, 0, 3}, {0, 0, 0}}

		require.Equal(t, expected, threeSum(input))
	})
}
