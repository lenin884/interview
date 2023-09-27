package rotate_list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("0->1->2, 4", func(t *testing.T) {
		head := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
		wantList := []int{2, 0, 1}
		got := rotateRight(head, 4)

		gotList := make([]int, 0)
		for got != nil {
			gotList = append(gotList, got.Val)
			got = got.Next
		}
		require.Equal(t, wantList, gotList)
	})
}
