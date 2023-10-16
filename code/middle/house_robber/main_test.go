package house_robber

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rob(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		got := rob(nums)
		want := 4

		require.Equal(t, want, got)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{2, 7, 9, 3, 1}
		got := rob(nums)
		want := 12

		require.Equal(t, want, got)
	})
}
