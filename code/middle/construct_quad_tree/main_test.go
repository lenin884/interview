package construct_quad_tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_construct(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		grid := [][]int{
			{1, 1},
			{1, 1},
		}
		expected := &Node{
			Val:    true,
			IsLeaf: true,
		}
		actual := construct(grid)
		require.Equal(t, expected, actual)
	})

	t.Run("case 2, large grid", func(t *testing.T) {
		grid := [][]int{
			{1, 1, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 0},
		}
		expectedValues := []bool{
			false,
			true,
			true,
			false,
		}

		actual := construct(grid)
		checkValues(t, actual, expectedValues)
	})
}

func checkValues(t *testing.T, node *Node, expected []bool) {
	if node == nil {
		return
	}
	if node.IsLeaf {
		require.Equal(t, expected[0], node.Val)
		return
	}
	checkValues(t, node.TopLeft, expected[1:])
	checkValues(t, node.TopRight, expected[2:])
	checkValues(t, node.BottomLeft, expected[3:])
	checkValues(t, node.BottomRight, expected[4:])
}
