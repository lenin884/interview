package sort_list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSortList(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		list := &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		}
		require.True(t, checkAscVal(sortList(list)))
	})

	t.Run("test getMiddle", func(t *testing.T) {
		list := &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
								Next: &ListNode{
									Val: 7,
								},
							},
						},
					},
				},
			},
		}
		require.Equal(t, 3, getMiddle(list).Val)
	})
}

func checkAscVal(node *ListNode) bool {
	if node == nil {
		return true
	}
	if node.Next == nil {
		return true
	}
	if node.Val > node.Next.Val {
		return false
	}
	return checkAscVal(node.Next)
}
