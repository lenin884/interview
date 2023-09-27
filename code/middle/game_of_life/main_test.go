package game_of_life

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		board := [][]int{
			{0, 1, 0},
			{0, 0, 1},
			{1, 1, 1}}
		want := [][]int{
			{0, 0, 0},
			{1, 0, 1},
			{0, 1, 1}}
		gameOfLife(board)
		require.Equal(t, want, board)
	})

	t.Run("test 2", func(t *testing.T) {
		board := [][]int{
			{1, 1},
			{1, 0}}
		want := [][]int{
			{1, 1},
			{1, 1}}
		gameOfLife(board)
		require.Equal(t, want, board)
	})

	t.Run("test 1 gameOfLifeNeetCode", func(t *testing.T) {
		board := [][]int{
			{0, 1, 0},
			{0, 0, 1},
			{1, 1, 1}}
		want := [][]int{
			{0, 0, 0},
			{1, 0, 1},
			{0, 1, 1}}
		gameOfLifeNeetCode(board)
		require.Equal(t, want, board)
	})

	t.Run("test 2 gameOfLifeNeetCode", func(t *testing.T) {
		board := [][]int{
			{1, 1},
			{1, 0}}
		want := [][]int{
			{1, 1},
			{1, 1}}
		gameOfLifeNeetCode(board)
		require.Equal(t, want, board)
	})
}
