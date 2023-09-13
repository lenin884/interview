package remove_duplicate_from_sorted_array_2

import "testing"

func TestName(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}

	got := removeDuplicates(nums)
	want := 5

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
