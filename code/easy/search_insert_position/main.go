package search_insert_position

func searchInsert(nums []int, target int) int {
	return searchInsertBinarySearch(nums, target)
}

// searchInsertBinarySearch time complexity O(log(n)), space complexity O(1)
func searchInsertBinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		}

		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
