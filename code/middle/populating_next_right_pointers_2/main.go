package populating_next_right_pointers_2

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}

	// Обрабатываем случай когда корень имеет соседа (next)
	if root.Next != nil {
		if root.Right != nil {
			root.Right.Next = findNext(root.Next)
		} else if root.Left != nil {
			root.Left.Next = findNext(root.Next)
		}
	}

	// сначала заполняем правые поддеревья чтобы были указатели справа
	connect(root.Right)
	// затем левые поддеревья
	connect(root.Left)

	return root
}

func findNext(node *Node) *Node {
	// Поиск следующего узла справа
	if node == nil {
		return nil
	}

	if node.Left != nil {
		return node.Left
	}

	if node.Right != nil {
		return node.Right
	}

	return findNext(node.Next)
}
