package longest_palindromic_substring

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_find(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := longestPalindrome("babad")
		require.Contains(t, []string{"bab", "aba"}, got)
	})

	t.Run("case 2", func(t *testing.T) {
		got := longestPalindrome("cbbd")
		want := "bb"
		require.Equal(t, want, got)
	})

	t.Run("case 3", func(t *testing.T) {
		got := longestPalindrome2("babad")
		require.Contains(t, []string{"bab", "aba"}, got)
	})

	t.Run("case 4", func(t *testing.T) {
		got := longestPalindrome2("cbbd")
		want := "bb"
		require.Equal(t, want, got)
	})
}
