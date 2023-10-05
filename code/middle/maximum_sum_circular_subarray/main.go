package maximum_sum_circular_subarray

func maxSubarraySumCircular(nums []int) int {
	globalMax, globalMin := nums[0], nums[0]
	localMax, localMin := 0, 0
	total := 0
	for i := 0; i < len(nums); i++ {
		localMax, localMin = max(localMax+nums[i], nums[i]), min(localMin+nums[i], nums[i])
		globalMax, globalMin = max(globalMax, localMax), min(globalMin, localMin)
		total += nums[i]
	}

	if globalMin == total { // when all number is negative
		return globalMax
	}
	return max(globalMax, total-globalMin)
}
