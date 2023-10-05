package implement_trie_prefix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTrie(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("apple")
		require.True(t, trie.Search("apple"))
		require.False(t, trie.Search("app"))
		require.True(t, trie.StartsWith("app"))
		trie.Insert("app")
		require.True(t, trie.Search("app"))
		trie.Insert("crew")
		require.True(t, trie.Search("crew"))
		require.False(t, trie.Search("crews"))
		require.True(t, trie.StartsWith("cre"))
	})
}
