package validate_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(node *TreeNode, min *TreeNode, max *TreeNode) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Val <= min.Val {
		return false
	}

	if max != nil && node.Val >= max.Val {
		return false
	}

	return validate(node.Left, min, node) && validate(node.Right, node, max)
}
