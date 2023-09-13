package minimum_absolute_difference_in_BST

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	return getMinimumDifferenceInorder(root)
}

// getMinimumDifferenceInorder time complexity O(n), space complexity O(1) - Morris algorithm
func getMinimumDifferenceInorder(root *TreeNode) int {
	var prev *TreeNode
	minDiff := math.MaxInt64

	// Morris algorithm
	// https://www.youtube.com/watch?v=wGXB9OWhPTg
	for root != nil {
		if root.Left == nil {
			if prev != nil {
				diff := root.Val - prev.Val
				if diff < minDiff {
					minDiff = diff
				}
			}
			prev = root
			root = root.Right
		} else {
			pred := root.Left
			// find predecessor
			for pred.Right != nil && pred.Right != root {
				pred = pred.Right
			}

			if pred.Right == nil {
				pred.Right = root
				root = root.Left
			} else {
				pred.Right = nil
				if prev != nil {
					diff := root.Val - prev.Val
					if diff < minDiff {
						minDiff = diff
					}
				}
				prev = root
				root = root.Right
			}
		}
	}

	return minDiff
}
