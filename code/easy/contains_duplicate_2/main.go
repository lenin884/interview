package contains_duplicate_2

func containsNearbyDuplicate(nums []int, k int) bool {
	numsIndex := make(map[int]int, 0)

	for i, num := range nums {
		if index, ok := numsIndex[num]; ok {
			if abs(i-index) <= k {
				return true
			}
		}

		numsIndex[num] = i
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
