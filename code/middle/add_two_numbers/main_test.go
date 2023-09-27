package add_two_numbers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
		l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
		result := addTwoNumbers(l1, l2)
		require.Equal(t, 7, result.Val)
		require.Equal(t, 0, result.Next.Val)
		require.Equal(t, 8, result.Next.Next.Val)
	})
}
