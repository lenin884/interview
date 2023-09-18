package reverse_words_in_string

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		input := "the sky is blue"
		excepted := "blue is sky the"

		require.Equal(t, excepted, reverseWords(input))
	})

	t.Run("test 2", func(t *testing.T) {
		input := "   hello    worlds   "
		excepted := "worlds hello"

		require.Equal(t, excepted, reverseWords(input))
	})
}
