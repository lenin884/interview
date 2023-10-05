package design_add_and_search_words_data_structure

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWordDictionary(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		wd := Constructor()
		wd.AddWord("bad")
		wd.AddWord("dad")
		wd.AddWord("mad")
		require.False(t, wd.Search("pad"))
		require.True(t, wd.Search("bad"))
		require.True(t, wd.Search(".ad"))
		require.True(t, wd.Search("b.."))
		require.True(t, wd.Search("..."))
	})
}
