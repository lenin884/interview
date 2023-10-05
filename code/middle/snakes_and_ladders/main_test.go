package snakes_and_ladders

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSnakesAndLadders(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		board := [][]int{
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 35, -1, -1, 13, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 15, -1, -1, -1, -1},
		}

		actual := snakesAndLadders(board)

		require.Equal(t, 4, actual)
	})

	t.Run("case 2 return -1", func(t *testing.T) {
		// [[1,1,-1],[1,1,1],[-1,1,1]]
		board := [][]int{
			{1, 1, -1},
			{1, 1, 1},
			{-1, 1, 1},
		}

		actual := snakesAndLadders(board)

		require.Equal(t, -1, actual)
	})
}
