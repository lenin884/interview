package roman_to_integer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	// Test function romanToInt. Input roman string and expected integer.
	t.Run("test 1", func(t *testing.T) {
		input := "III"
		expected := 3

		require.Equal(t, expected, romanToInt(input))
	})

	t.Run("test 2", func(t *testing.T) {
		input := "IV"
		expected := 4

		require.Equal(t, expected, romanToInt(input))
	})

	t.Run("complex number", func(t *testing.T) {
		input := "MCMXCIV"
		expected := 1994

		require.Equal(t, expected, romanToInt(input))
	})
}
