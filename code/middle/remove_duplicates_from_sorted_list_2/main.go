package remove_duplicates_from_sorted_list_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	hashmap := make(map[int]int)

	curr := head
	for curr != nil {
		hashmap[curr.Val]++
		curr = curr.Next
	}

	curr = head
	var prev *ListNode
	for curr != nil {
		if hashmap[curr.Val] > 1 {
			if prev == nil {
				head = curr.Next
			} else {
				prev.Next = curr.Next
			}
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	return head
}
