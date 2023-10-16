package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) / 2)
		if target == nums[mid] {
			return mid
		}

		// left side is sorted
		if nums[left] <= nums[mid] {
			// target is in the left side
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				// target is in the right side
				left = mid + 1
			}
		} else {
			// right side is sorted
			// target is in the right side
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				// target is in the left side
				right = mid - 1
			}
		}
	}
	return -1
}
