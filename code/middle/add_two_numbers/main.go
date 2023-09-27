package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
python code:
def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:

	curr = ListNode(0)
	curr_head = curr
	carry = 0
	while l1 or l2 or carry:
	    x = l1.val if l1 else 0
	    y = l2.val if l2 else 0
	    sum_val = carry + x + y
	    carry, node_val = divmod(sum_val, 10)
	    curr_head.next = ListNode(node_val)
	    curr_head = curr_head.next
	    l1 = l1.next if l1 else None
	    l2 = l2.next if l2 else None

	return curr.next
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr := &ListNode{}
	currHead := curr
	carry := 0
	nodeVal := 0
	for l1 != nil || l2 != nil || carry != 0 {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sumVal := carry + x + y
		carry, nodeVal = sumVal/10, sumVal%10
		currHead.Next = &ListNode{Val: nodeVal}
		currHead = currHead.Next
	}

	return curr.Next
}
