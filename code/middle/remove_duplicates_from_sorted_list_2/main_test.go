package remove_duplicates_from_sorted_list_2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	t.Run("1->2->3->3->4->4->5", func(t *testing.T) {
		head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3,
			Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4,
				Next: &ListNode{Val: 5}}}}}}}
		wantList := []int{1, 2, 5}
		got := deleteDuplicates(head)

		gotList := make([]int, 0)
		for got != nil {
			gotList = append(gotList, got.Val)
			got = got.Next
		}
		require.Equal(t, wantList, gotList)
	})
}
