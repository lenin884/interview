package word_break

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		require.Equal(t, true, wordBreak("leetcode", []string{"leet", "code"}))
	})

	t.Run("case 2", func(t *testing.T) {
		require.Equal(t, true, wordBreak("applepenapple", []string{"apple", "pen"}))
	})

	t.Run("case 3", func(t *testing.T) {
		require.Equal(t, false, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	})
}
