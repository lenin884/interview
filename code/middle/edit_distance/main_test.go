package edit_distance

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_minDistance(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		word1 := "horse"
		word2 := "ros"
		require.Equal(t, 3, minDistance(word1, word2))
	})

	t.Run("case 2", func(t *testing.T) {
		word1 := "intention"
		word2 := "execution"
		require.Equal(t, 5, minDistance(word1, word2))
	})
}
