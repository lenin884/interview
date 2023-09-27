package simplify_path

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		got := simplifyPath("/home/")
		want := "/home"
		require.Equal(t, want, got)
	})

	t.Run("test two", func(t *testing.T) {
		got := simplifyPath("/../")
		want := "/"
		require.Equal(t, want, got)
	})

	t.Run("test three", func(t *testing.T) {
		got := simplifyPath("/home//foo/")
		want := "/home/foo"
		require.Equal(t, want, got)
	})
}
