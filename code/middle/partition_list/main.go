package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	/*
			Algorithm with one pass:
			1. Create two pointers, one for less than x, one for greater than or equal to x
			2. Iterate through the list, if the current node is less than x, add it to the less than list, otherwise add it to the greater than or equal to list
		    3. At the end, connect the two lists together
	*/
	lessThanHead := &ListNode{}
	lessThanTail := lessThanHead

	greaterThanHead := &ListNode{}
	greaterThanTail := greaterThanHead

	for head != nil {
		if head.Val < x {
			lessThanTail.Next = head
			lessThanTail = lessThanTail.Next
		} else {
			greaterThanTail.Next = head
			greaterThanTail = greaterThanTail.Next
		}
		head = head.Next
	}

	lessThanTail.Next = greaterThanHead.Next
	greaterThanTail.Next = nil

	return lessThanHead.Next
}
