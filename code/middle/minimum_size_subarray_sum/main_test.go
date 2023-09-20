package minimum_size_subarray_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		input := []int{2, 3, 1, 2, 4, 3}
		target := 7
		expected := 2

		require.Equal(t, expected, minSubArrayLen(target, input))
	})

	t.Run("test 2", func(t *testing.T) {
		input := []int{1, 4, 4}
		target := 4
		expected := 1

		require.Equal(t, expected, minSubArrayLen(target, input))
	})
}
