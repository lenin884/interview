package remove_duplicate_from_sorted_array_2

func removeDuplicates(nums []int) int {
	// edge case
	if len(nums) < 3 {
		return len(nums)
	}

	// slow pointer
	slow := 2
	// fast pointer
	for fast := 2; fast < len(nums); fast++ {
		// 1 1 1 2 2 3
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
