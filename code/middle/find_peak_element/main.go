package find_peak_element

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + ((right - left) / 2)
		// left neighbor is greater
		if mid > 0 && nums[mid-1] > nums[mid] {
			right = mid - 1
		} else if mid < len(nums)-1 && nums[mid+1] > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
