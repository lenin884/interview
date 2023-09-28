package populating_next_right_pointers_2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConnect(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		root := &Node{
			Val: 1,
			Left: &Node{
				Val: 2,
				Left: &Node{
					Val: 4,
					Left: &Node{
						Val: 7,
					},
				},
				Right: &Node{
					Val: 5,
				},
			},
			Right: &Node{
				Val: 3,
				Right: &Node{
					Val: 6,
					Right: &Node{
						Val: 8,
					},
				},
			},
		}
		connect(root)
		// check next pointers
		require.Nil(t, root.Next)
		require.Equal(t, 3, root.Left.Next.Val)
		require.Nil(t, root.Right.Next)
		require.Equal(t, 5, root.Left.Left.Next.Val)
		require.Equal(t, 6, root.Left.Right.Next.Val)
		require.Equal(t, 8, root.Left.Left.Left.Next.Val)

	})
}
