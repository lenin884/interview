package find_first_and_last_position_in_sorted_array

func searchRange(nums []int, target int) []int {
	ranges := []int{-1, -1}
	if len(nums) == 0 {
		return ranges
	}

	ranges[0] = searchRangeBias(nums, target, true)
	ranges[1] = searchRangeBias(nums, target, false)

	return ranges
}

func searchRangeBias(nums []int, target int, leftBias bool) int {
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			index = mid
			if leftBias {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return index
}
