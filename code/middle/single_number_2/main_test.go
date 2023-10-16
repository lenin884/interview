package single_number_2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{2, 2, 3, 2}
		got := singleNumber(nums)
		want := 3
		require.Equal(t, want, got)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{0, 1, 0, 1, 0, 1, 99, 3, 2, 3, 2, 3, 2}
		got := singleNumber(nums)
		want := 99
		require.Equal(t, want, got)
	})
}
