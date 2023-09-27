package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
flatten converts a binary tree to a linked list. On O(1) space.
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// Flatten left subtree
	flatten(root.Left)

	// Flatten right subtree
	flatten(root.Right)

	// Move left subtree to right
	if root.Left != nil {
		// Find the rightmost node of left subtree
		rightmost := root.Left
		for rightmost.Right != nil {
			rightmost = rightmost.Right
		}

		// Move right subtree to the rightmost node of left subtree
		rightmost.Right = root.Right

		// Move left subtree to right
		root.Right = root.Left
		root.Left = nil
	}
}
