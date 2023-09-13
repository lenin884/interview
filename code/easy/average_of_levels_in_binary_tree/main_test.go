package average_of_levels_in_binary_tree

import "testing"

func TestAverage(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	got := averageOfLevelsBFS(root)
	want := []float64{3, 14.5, 11}

	if len(got) != len(want) {
		t.Errorf("got %v want %v", got, want)
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
