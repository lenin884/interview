package generate_parentheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_generateParentheses(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		require.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParentheses(3))
	})
}
