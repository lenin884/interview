package letter_combination_phone_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		require.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	})

	t.Run("case 2", func(t *testing.T) {
		require.Equal(t, []string{}, letterCombinations(""))
	})

	t.Run("case 3", func(t *testing.T) {
		require.Equal(t, []string{"a", "b", "c"}, letterCombinations("2"))
	})

	t.Run("case 4", func(t *testing.T) {
		digits := "234"
		expected := []string{
			"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
			"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
			"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
		}

		require.Equal(t, expected, letterCombinations(digits))
	})
}
