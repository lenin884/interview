package triangle

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_triangle(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		triangle := [][]int{
			{2},
			{3, 4},
			{6, 5, 7},
			{4, 1, 8, 3},
		}

		got := minimumTotal(triangle)
		want := 11

		require.Equal(t, want, got)
	})
}
