package minimum_absolute_difference_in_BST

import "testing"

func TestName(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: -2,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: -1,
					},
				},
			},
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		Right: &TreeNode{
			Val: 30,
			Right: &TreeNode{
				Val: 40,
			},
		},
	}

	got := getMinimumDifference(root)
	want := 1

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
