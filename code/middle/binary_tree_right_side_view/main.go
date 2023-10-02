package binary_tree_right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {

	// реализуем алгоритм с помощью bfs
	// нужно пройти по всем уровням и взять последний элемент

	if root == nil {
		return []int{}
	}

	res := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if i == size-1 {
				res = append(res, curr.Val)
			}

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}

	return res
}
