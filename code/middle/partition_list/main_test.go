package partition_list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		input := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  5,
							Next: &ListNode{Val: 2},
						},
					},
				},
			},
		}
		expected := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}},
					},
				},
			},
		}
		actual := partition(input, 3)
		require.Condition(t, func() bool {
			actualNode := actual
			expectedNode := expected

			for actualNode != nil && expectedNode != nil {
				if actualNode.Val != expectedNode.Val {
					return false
				}
				actualNode = actualNode.Next
				expectedNode = expectedNode.Next
			}

			if (actualNode != nil && expectedNode != nil) && (actualNode == nil || expectedNode == nil) {
				return false
			}

			return true
		})
	})
}
