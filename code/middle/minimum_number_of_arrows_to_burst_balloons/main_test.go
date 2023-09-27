package minimum_number_of_arrows_to_burst_balloons

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}

		require.Equal(t, 2, findMinArrowShots(points))
	})

	t.Run("test two", func(t *testing.T) {
		points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}

		require.Equal(t, 4, findMinArrowShots(points))
	})

	t.Run("test three", func(t *testing.T) {
		points := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}

		require.Equal(t, 2, findMinArrowShots(points))
	})

	t.Run("test four", func(t *testing.T) {
		points := [][]int{{1, 20}, {2, 15}, {7, 12}, {10, 16}}

		require.Equal(t, 1, findMinArrowShots(points))
	})
}
