package same_tree

import "testing"

func TestName(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	if !isSameTree1(p, q) {
		t.Errorf("\nRecursive\ninput: %v\nexpected: %t\nactual: %t\n", p, true, false)
	}

	if !isSameTree2(p, q) {
		t.Errorf("\nQueue\ninput: %v\nexpected: %t\nactual: %t\n", p, true, false)
	}
}
