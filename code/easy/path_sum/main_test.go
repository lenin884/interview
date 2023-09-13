package path_sum

import "testing"

func TestHasPathSum1(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	if hasPathSum1(p, 5) {
		t.Errorf("\nRecursive\ninput: %v\nexpected: %t\nactual: %t\n", p, false, true)
	}

	if !hasPathSum1(p, 3) {
		t.Errorf("\nRecursive\ninput: %v\nexpected: %t\nactual: %t\n", p, true, false)
	}
}

func TestHasPathSum2(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	if hasPathSum2(p, 5) {
		t.Errorf("\nQueue\ninput: %v\nexpected: %t\nactual: %t\n", p, false, true)
	}

	if !hasPathSum2(p, 3) {
		t.Errorf("\nQueue\ninput: %v\nexpected: %t\nactual: %t\n", p, true, false)
	}
}
