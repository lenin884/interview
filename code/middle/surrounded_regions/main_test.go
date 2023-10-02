package surrounded_regions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_solve(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		board := [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}
		expected := [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}
		solve(board)
		require.Equal(t, expected, board)
	})

	t.Run("case 1", func(t *testing.T) {
		board := [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}
		expected := [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}
		solve(board)
		require.Equal(t, expected, board)
	})
}
