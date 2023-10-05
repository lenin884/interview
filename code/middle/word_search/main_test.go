package word_search

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_exist(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		board := [][]byte{
			{'D', 'A', 'B', 'C', 'E'},
			{'S', 'S', 'F', 'C', 'S'},
			{'C', 'A', 'D', 'E', 'E'},
		}
		word := "ABCCED"
		got := exist(board, word)
		require.True(t, got)
	})
}
