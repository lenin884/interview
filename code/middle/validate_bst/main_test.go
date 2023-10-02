package validate_bst

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	t.Run("should return true for valid BST", func(t *testing.T) {
		root := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}

		require.True(t, isValidBST(root))
	})

	t.Run("should return false for invalid BST", func(t *testing.T) {
		root := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		}

		require.False(t, isValidBST(root))
	})
}
