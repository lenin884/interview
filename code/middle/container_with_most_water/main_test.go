package container_with_most_water

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		expected := 49

		require.Equal(t, expected, maxArea(input))
	})

	t.Run("test 2", func(t *testing.T) {
		input := []int{1, 1}
		expected := 1

		require.Equal(t, expected, maxArea(input))
	})

	t.Run("test 3", func(t *testing.T) {
		input := []int{4, 3, 2, 1, 4}
		expected := 16

		require.Equal(t, expected, maxArea(input))
	})
}
