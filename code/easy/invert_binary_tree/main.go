package invert_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree1 is a recursive solution
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree1(root.Right), invertTree1(root.Left)

	return root
}

// invertTree2 is a iterative solution
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}
