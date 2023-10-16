package longest_increasing_subsequence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
		expect := 4
		require.Equal(t, expect, lengthOfLIS(nums))
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{0, 1, 0, 3, 2, 3}
		expect := 4
		require.Equal(t, expect, lengthOfLIS(nums))
	})

	t.Run("case 3", func(t *testing.T) {
		nums := []int{7, 7, 7, 7, 7, 7, 7}
		expect := 1
		require.Equal(t, expect, lengthOfLIS(nums))
	})

	t.Run("case 4", func(t *testing.T) {
		nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
		expect := 6
		require.Equal(t, expect, lengthOfLIS(nums))
	})
}
