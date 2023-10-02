package binary_tree_zigzag_level_order_traversal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	t.Run("should return [[3],[20,9],[15,7]] when input [3,9,20,null,null,15,7]", func(t *testing.T) {
		root := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}

		expected := [][]int{
			{3},
			{20, 9},
			{15, 7},
		}

		require.Equal(t, expected, zigzagLevelOrder(root))
	})
}
