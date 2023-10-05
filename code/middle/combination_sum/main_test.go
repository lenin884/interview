package combination_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := []int{2, 3, 6, 7}
		target := 7
		expected := [][]int{{2, 2, 3}, {7}}
		actual := combinationSum(input, target)
		require.Equal(t, expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		/*
			[8,7,4,3]
			11
		*/
		input := []int{8, 7, 4, 3}
		target := 11
		expected := [][]int{{3, 4, 4}, {3, 8}, {4, 7}}
		actual := combinationSum(input, target)
		require.Equal(t, expected, actual)
	})
}
