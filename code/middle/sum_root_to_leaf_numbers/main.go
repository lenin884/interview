package sum_root_to_leaf_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum = sum*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return sum
	}

	return dfs(root.Left, sum) + dfs(root.Right, sum)
}
