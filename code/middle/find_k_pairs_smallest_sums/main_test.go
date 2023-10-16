package find_k_pairs_smallest_sums

import (
	"testing"
)

func Test_kSmallestSum(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums1 := []int{1, 7, 11}
		nums2 := []int{2, 4, 6}
		k := 3
		got := kSmallestPairs(nums1, nums2, k)
		want := [][]int{{1, 2}, {1, 4}, {1, 6}}

		if !checkContains(want, got) {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}

func checkContains(actual, want [][]int) bool {
	contain := false
	for _, w := range want {
		for _, a := range actual {
			if w[0] == a[0] && w[1] == a[1] {
				contain = true
				break
			}
		}

		if !contain {
			return false
		}
		contain = false
	}

	return true
}
