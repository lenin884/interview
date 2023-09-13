package count_complete_tree_nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	return countNodesRecursive(root)
}

func countNodesRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + countNodesRecursive(root.Left) + countNodesRecursive(root.Right)
}
