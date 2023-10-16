package kth_largest_element_in_array

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	return quickSelect(0, len(nums)-1, k, nums)
}

func quickSelect(left, right, k int, nums []int) int {
	pivot, p := nums[right], left
	for i := left; i < right; i++ {
		if nums[i] <= pivot {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	nums[p], nums[right] = nums[right], nums[p]

	if p > k {
		return quickSelect(left, p-1, k, nums)
	} else if p < k {
		return quickSelect(p+1, right, k, nums)
	} else {
		return nums[p]
	}
}
