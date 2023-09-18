package int_to_roman

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	// Test function romanToInt. Input roman string and expected integer.
	t.Run("test 1", func(t *testing.T) {
		input := 3
		expected := "III"

		require.Equal(t, expected, intToRoman(input))
	})

	t.Run("test 2", func(t *testing.T) {
		input := 4
		expected := "IV"

		require.Equal(t, expected, intToRoman(input))
	})

	t.Run("complex number", func(t *testing.T) {
		input := 1994
		expected := "MCMXCIV"

		require.Equal(t, expected, intToRoman(input))
	})

	t.Run("test 3", func(t *testing.T) {
		input := 78
		expected := "LXXVIII"

		require.Equal(t, expected, intToRoman(input))
	})
}
