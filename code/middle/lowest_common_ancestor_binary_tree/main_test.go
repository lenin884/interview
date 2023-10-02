package lowest_common_ancestor_binary_tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLowest(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		}

		p := &TreeNode{
			Val: 5,
		}

		q := &TreeNode{
			Val: 4,
		}

		require.Equal(t, 5, lowestCommonAncestor(root, p, q).Val)
	})

	t.Run("test 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}

		p := &TreeNode{
			Val: 2,
		}

		q := &TreeNode{
			Val: 3,
		}

		require.Equal(t, 1, lowestCommonAncestor(root, p, q).Val)
	})
}
