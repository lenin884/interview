package group_anagram

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		inputs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		expected := [][]string{
			{"ate", "eat", "tea"},
			{"nat", "tan"},
			{"bat"},
		}

		require.Equal(t, expected, groupAnagrams(inputs))
	})

	t.Run("test 2 best", func(t *testing.T) {
		inputs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		expected := [][]string{
			{"ate", "eat", "tea"},
			{"nat", "tan"},
			{"bat"},
		}

		require.Equal(t, expected, groupAnagramsBest(inputs))
	})
}
