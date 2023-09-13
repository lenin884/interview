package count_complete_tree_nodes

import "testing"

func TestRecursive(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	got := countNodesRecursive(root)
	want := 2

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
