package construct_binary_from_preorder_and_inorder_travelsal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuildTree(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		preorder := []int{1, 2}
		inorder := []int{2, 1}
		root := buildTree(preorder, inorder)
		require.Equal(t, 1, root.Val)
		require.Equal(t, 2, root.Left.Val)
		require.Nil(t, root.Right)
	})

	t.Run("test two", func(t *testing.T) {
		preorder := []int{3, 9, 20, 15, 7}
		inorder := []int{9, 3, 15, 20, 7}
		root := buildTree(preorder, inorder)
		require.Equal(t, 3, root.Val)
		require.Equal(t, 9, root.Left.Val)
		require.Equal(t, 20, root.Right.Val)
		require.Equal(t, 15, root.Right.Left.Val)
		require.Equal(t, 7, root.Right.Right.Val)
	})
}
