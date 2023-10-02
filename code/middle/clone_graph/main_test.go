package clone_graph

import (
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		node1 := &Node{Val: 1}
		node2 := &Node{Val: 2}
		node3 := &Node{Val: 3}
		node4 := &Node{Val: 4}
		node1.Neighbors = []*Node{node2, node4}
		node2.Neighbors = []*Node{node1, node3}
		node3.Neighbors = []*Node{node2, node4}
		node4.Neighbors = []*Node{node1, node3}
		expected := &Node{Val: 1}
		expected.Neighbors = []*Node{
			&Node{Val: 2},
			&Node{Val: 4},
		}
		expected.Neighbors[0].Neighbors = []*Node{
			&Node{Val: 1},
			&Node{Val: 3},
		}
		expected.Neighbors[1].Neighbors = []*Node{
			&Node{Val: 1},
			&Node{Val: 3},
		}
		actual := cloneGraph(node1)
		t.Error(t, expected, actual)
	})
}
