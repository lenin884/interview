package summary_ranges

import "testing"

func TestName(t *testing.T) {
	nums := [][]int{{0, 1, 2, 4, 5, 7}, {0, 2, 3, 4, 6, 8, 9}, {}, {-1}, {0}}
	expected := [][]string{{"0->2", "4->5", "7"}, {"0", "2->4", "6", "8->9"}, {}, {"-1"}, {"0"}}
	for i := 0; i < len(nums); i++ {
		if !compare(summaryRanges(nums[i]), expected[i]) {
			t.Errorf("\ninput: %v\nexpected: %v\nactual: %v", nums[i], expected[i], summaryRanges(nums[i]))
		}
	}
}

func compare(result []string, expected []string) bool {
	if len(result) != len(expected) {
		return false
	}
	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			return false
		}
	}
	return true
}
