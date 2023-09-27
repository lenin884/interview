package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	length := 1

	tail := head
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	if length == 0 {
		return nil
	}

	k = k % length
	if k == 0 {
		return head
	}

	curr := head
	for i := 0; i < length-k-1; i++ {
		curr = curr.Next
	}

	newHead := curr.Next
	curr.Next = nil
	tail.Next = head

	return newHead
}
