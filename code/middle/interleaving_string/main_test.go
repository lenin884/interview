package interleaving_string

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		s1 := "aabcc"
		s2 := "dbbca"
		s3 := "aadbbcbcac"
		got := isInterleave(s1, s2, s3)
		require.Equal(t, true, got)
	})

	t.Run("case 2", func(t *testing.T) {
		s1 := "aabcc"
		s2 := "dbbca"
		s3 := "aadbbbaccc"
		got := isInterleave(s1, s2, s3)
		require.Equal(t, false, got)
	})
}
