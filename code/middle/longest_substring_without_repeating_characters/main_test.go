package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("abcabcbb", func(t *testing.T) {
		require.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	})

	t.Run("bbbbb", func(t *testing.T) {
		require.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	})

	t.Run("pwwkew", func(t *testing.T) {
		require.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	})
}
