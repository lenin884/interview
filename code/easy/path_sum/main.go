package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// hasPathSum1 is a recursive solution
func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && targetSum-root.Val == 0 {
		return true
	}

	return hasPathSum1(root.Left, targetSum-root.Val) || hasPathSum1(root.Right, targetSum-root.Val)
}

// hasPathSum2 is a iterative solution
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}
	sums := []int{root.Val}

	for len(queue) > 0 {
		node := queue[0]
		currSum := sums[0]
		queue = queue[1:]
		sums = sums[1:]

		// Check if it's a leaf node and if the sum equals the targetSum.
		if node.Left == nil && node.Right == nil && currSum == targetSum {
			return true
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			sums = append(sums, currSum+node.Left.Val)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
			sums = append(sums, currSum+node.Right.Val)
		}
	}

	return false
}
