package same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree1(p *TreeNode, q *TreeNode) bool {
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}

		if !isSameTree1(p.Left, q.Left) {
			return false
		}

		if !isSameTree1(p.Right, q.Right) {
			return false
		}

		return true
	}

	if p == nil && q == nil {
		return true
	}

	return false
}

// isSameTree2 is a using queue to solve the problem
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}

	for len(queue) > 0 {
		node1 := queue[0]
		node2 := queue[1]
		queue = queue[2:]

		// If both nodes are nil, continue
		if node1 == nil && node2 == nil {
			continue
		}

		// If one node is not nil, return false
		if node1 == nil || node2 == nil {
			return false
		}

		// If both nodes are not nil, but their values are not equal, return false
		if node1.Val != node2.Val {
			return false
		}

		queue = append(queue, node1.Left, node2.Left, node1.Right, node2.Right)
	}

	return true
}
