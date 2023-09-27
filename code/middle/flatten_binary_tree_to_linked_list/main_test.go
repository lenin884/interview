package flatten_binary_tree_to_linked_list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlatten(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		flatten(root)
		require.Equal(t, 1, root.Val)
		require.Equal(t, 2, root.Right.Val)
		require.Equal(t, 3, root.Right.Right.Val)
		require.Equal(t, 4, root.Right.Right.Right.Val)
		require.Equal(t, 5, root.Right.Right.Right.Right.Val)
		require.Equal(t, 6, root.Right.Right.Right.Right.Right.Val)
		require.Equal(t, 7, root.Right.Right.Right.Right.Right.Right.Val)
		require.Nil(t, root.Right.Right.Right.Right.Right.Right.Right)
	})
}
