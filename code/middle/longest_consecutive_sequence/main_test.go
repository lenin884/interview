package longest_consecutive_sequence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test longest consecutive sequence 1", func(t *testing.T) {
		nums := []int{100, 4, 200, 1, 3, 2}

		require.Equal(t, 4, longestConsecutive(nums))
	})

	t.Run("test longest consecutive sequence 2", func(t *testing.T) {
		nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

		require.Equal(t, 9, longestConsecutive(nums))
	})
}
