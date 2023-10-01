package binary_search_tree_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Реализуем алгоритм следующим образом:
Нужно хранить текущий элемент всегда левый нижний элемент, если его нет, то берем элемент выше по дереву и правый элемент

++++3++++
++/+++\++
++1++++5+
+/+\++/++
0+++4+6++

0 -> 1 -> 2 -> 4 -> 3 -> 6 -> 5 - inorder traversal

*/

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	// инициализация
	// сначала нужно найти самый левый элемент

	ti := BSTIterator{
		stack: []*TreeNode{},
	}

	for root != nil {
		ti.stack = append(ti.stack, root)
		root = root.Left
	}

	return ti
}

func (ti *BSTIterator) Next() int {
	// возвращает следующий элемент
	res := ti.stack[len(ti.stack)-1]
	cur := res.Right
	ti.stack = ti.stack[:len(ti.stack)-1]

	for cur != nil {
		ti.stack = append(ti.stack, cur)
		cur = cur.Left
	}

	return res.Val
}

func (ti *BSTIterator) HasNext() bool {
	// есть ли следующий элемент
	if len(ti.stack) > 0 {
		return true
	}
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
