package zigzag_conversion

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		input := "PAYPALISHIRING"
		expected := "PAHNAPLSIIGYIR"

		require.Equal(t, expected, convert(input, 3))
	})

	t.Run("test 2", func(t *testing.T) {
		input := "PAYPALISHIRING"
		expected := "PHASIYIRPLIGAN"

		require.Equal(t, expected, convert(input, 5))
	})

	t.Run("test 3", func(t *testing.T) {
		input := "ABC"
		expected := "ACB"

		require.Equal(t, expected, convert(input, 2))
	})
}
