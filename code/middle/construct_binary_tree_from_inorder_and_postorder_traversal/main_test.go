package construct_binary_tree_from_inorder_and_postorder_traversal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuildTree(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		inorder := []int{2, 1}
		postorder := []int{2, 1}
		root := buildTree(inorder, postorder)
		require.Equal(t, 1, root.Val)
		require.Equal(t, 2, root.Left.Val)
		require.Nil(t, root.Right)
	})

	t.Run("test two", func(t *testing.T) {
		inorder := []int{9, 3, 15, 20, 7}
		postorder := []int{9, 15, 7, 20, 3}
		root := buildTree(inorder, postorder)
		require.Equal(t, 3, root.Val)
		require.Equal(t, 9, root.Left.Val)
		require.Equal(t, 20, root.Right.Val)
		require.Equal(t, 15, root.Right.Left.Val)
		require.Equal(t, 7, root.Right.Right.Val)
	})
}
