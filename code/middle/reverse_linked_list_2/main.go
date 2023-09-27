package reverse_linked_list_2

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
reverseBetween reverses the nodes of the list from position left to position right, and returns the reversed list.
 1. Create a list of nodes from the given list.
 2. Reverse the list of nodes from position left to position right.
 3. Create a new list from the reversed list of nodes.
 4. Return the new list.
 5. Complexity:
    5.1 Time Complexity: O(n)
    5.2 Space Complexity: O(n)
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	list := make([]*ListNode, 0)

	for curr := head; curr != nil; curr = curr.Next {
		list = append(list, curr)
	}

	for i, j := left-1, right-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	for i := 0; i < len(list)-1; i++ {
		list[i].Next = list[i+1]
	}

	list[len(list)-1].Next = nil

	return list[0]
}
