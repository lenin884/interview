package bitwise_and_of_numbers_range

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Bitwise(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := rangeBitwiseAnd(5, 7)
		want := 4
		require.Equal(t, want, got)
	})

	t.Run("case 2", func(t *testing.T) {
		got := rangeBitwiseAnd(0, 1)
		want := 0
		require.Equal(t, want, got)
	})
}
