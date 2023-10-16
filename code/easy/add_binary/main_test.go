package add_binary

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_addBinary(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		a := "11"
		b := "1"
		got := addBinary(a, b)
		want := "100"
		require.Equal(t, want, got)
	})
}
