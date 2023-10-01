package binary_search_tree_iterator

import "testing"

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
				},
			},
		}

		ti := Constructor(root)
		t.Log(ti.Next())    // 3
		t.Log(ti.Next())    // 7
		t.Log(ti.HasNext()) // true
		t.Log(ti.Next())    // 9
		t.Log(ti.HasNext()) // true
		t.Log(ti.Next())    // 15
		t.Log(ti.HasNext()) // true
		t.Log(ti.Next())    // 20
		t.Log(ti.HasNext()) // false
	})
}
