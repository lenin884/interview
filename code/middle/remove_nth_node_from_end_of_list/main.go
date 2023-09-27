package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// resolve this by two pointer (shift window)
	left := head
	right := head

	// shift right pointer to n
	for right != nil && n > 0 {
		right = right.Next
		n--
	}

	// if right is nil, it means n is equal to length of list
	// so we need to remove head
	if right == nil {
		return head.Next
	}

	// shift left and right pointer until right is nil
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	// remove nth node
	left.Next = left.Next.Next

	return head
}
