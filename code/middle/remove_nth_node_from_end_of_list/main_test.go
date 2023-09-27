package remove_nth_node_from_end_of_list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("should return [1,2,3,5]", func(t *testing.T) {
		input := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
		n := 2
		actual := removeNthFromEnd(input, n)

		expected := []int{1, 2, 3, 5}
		actualList := make([]int, 0)
		for actual != nil {
			actualList = append(actualList, actual.Val)
			actual = actual.Next
		}

		require.Equal(t, expected, actualList)
	})
}
