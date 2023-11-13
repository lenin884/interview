package text_justification

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJustify(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		words := []string{"This", "is", "an", "example", "of", "text", "justification."}
		maxWidth := 16
		expected := []string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		}

		require.Equalf(t, expected, fullJustify(words, maxWidth), "fullJustify(%v, %d) should be %v", words, maxWidth, expected)
	})

	t.Run("case 2", func(t *testing.T) {
		words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
		maxWidth := 16
		expected := []string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		}

		require.Equalf(t, expected, fullJustify(words, maxWidth), "fullJustify(%v, %d) should be %v", words, maxWidth, expected)
	})
}
