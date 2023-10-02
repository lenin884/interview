package kth_smallest_element_in_bst

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		}
		k := 1
		require.Equal(t, 1, kthSmallest(root, k))
	})

	t.Run("case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 6,
			},
		}
		k := 3
		require.Equal(t, 3, kthSmallest(root, k))
	})
}
