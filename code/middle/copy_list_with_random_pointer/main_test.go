package copy_list_with_random_pointer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var head *Node
		require.Nil(t, copyRandomList(head))
	})

	t.Run("test 1", func(t *testing.T) {
		head := &Node{
			Val: 1,
		}
		head.Next = &Node{
			Val: 2,
		}
		head.Next.Next = &Node{
			Val: 3,
		}
		head.Next.Next.Next = &Node{
			Val: 4,
		}
		head.Next.Next.Next.Next = &Node{
			Val: 5,
		}
		head.Random = head.Next.Next
		head.Next.Random = head.Next.Next.Next.Next
		head.Next.Next.Next.Random = head.Next

		newHead := copyRandomList(head)
		require.Equal(t, 1, newHead.Val)
		require.Equal(t, 2, newHead.Next.Val)
		require.Equal(t, 3, newHead.Next.Next.Val)
		require.Equal(t, 4, newHead.Next.Next.Next.Val)
		require.Equal(t, 5, newHead.Next.Next.Next.Next.Val)
		require.Equal(t, 3, newHead.Random.Val)
		require.Equal(t, 5, newHead.Next.Random.Val)
		require.Equal(t, 2, newHead.Next.Next.Next.Random.Val)
		require.Nil(t, newHead.Next.Next.Next.Next.Random)
	})
}
