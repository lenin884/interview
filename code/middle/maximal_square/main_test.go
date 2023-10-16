package maximal_square

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		matrix := [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}
		expect := 4
		require.Equal(t, expect, maximalSquare(matrix))
	})

	t.Run("case 2", func(t *testing.T) {
		matrix := [][]byte{
			{'0', '1'},
			{'1', '0'},
		}
		expect := 1
		require.Equal(t, expect, maximalSquare(matrix))
	})
}
