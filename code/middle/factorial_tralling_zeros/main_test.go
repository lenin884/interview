package factorial_tralling_zeros

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := trailingZeroes(3)
		require.Equal(t, 0, got)
	})

	t.Run("case 2", func(t *testing.T) {
		got := trailingZeroes(5)
		require.Equal(t, 1, got)
	})
}
