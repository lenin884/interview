package average_of_levels_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	return averageOfLevelsBFS(root)
}

func averageOfLevelsBFS(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	var result []float64

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		var sum float64
		var count int

		for i := len(stack); i > 0; i-- {
			node := stack[0]
			stack = stack[1:]

			sum += float64(node.Val)
			count++

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}

		result = append(result, sum/float64(count))
	}

	return result
}
