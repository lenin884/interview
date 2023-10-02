package lowest_common_ancestor_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// First base case, if root is null, return null
	if root == nil {
		return nil
	}

	// Second base case, found our case
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// get left and then right nodes
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// if both left and right aren't null, that means we found the targets on both sides of trees, means we need to return root
	if left != nil && right != nil {
		return root
	}

	// if we find in left, return left
	if left != nil {
		return left
	}

	// else right
	return right
}
