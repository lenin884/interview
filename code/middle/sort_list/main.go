package sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // empty list or list with one element
		return head
	}

	// split list into two parts
	left := head
	right := getMiddle(head)
	tmp := right.Next
	right.Next = nil
	right = tmp

	// sort each part
	left = sortList(left)     // left part is sorted
	right = sortList(right)   // right part is sorted
	return merge(left, right) // merge two sorted parts
}

// нужно посмотреть отдельно на списке
func getMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(left, right *ListNode) *ListNode {
	tail := &ListNode{}
	dummy := tail

	for left != nil && right != nil {
		if left.Val < right.Val {
			tail.Next = left
			left = left.Next
		} else {
			tail.Next = right
			right = right.Next
		}
		tail = tail.Next
	}

	if left != nil {
		tail.Next = left
	}
	if right != nil {
		tail.Next = right
	}

	return dummy.Next
}
