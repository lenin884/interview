package search_insert_position

import "testing"

func TestName(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2

	got := searchInsert(nums, target)
	want := 1

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
