package minimum_size_subarray_sum

import "math"

func minSubArrayLen(target int, nums []int) int {
	/*
		Implementation of the algorithm:
		1. Initialize left pointer to 0 and sum to 0
		2. Iterate over the nums:
			2.1. Add nums[i] to sum
			2.2. While sum >= target:
				2.2.1. Update the result if the current subarray's length is less than the minimum found so far.
				2.2.2. Subtract nums[left] from sum and increment left.
		3. Return result
	*/
	left, total := 0, 0
	result := math.MaxInt

	for right := 0; right < len(nums); right++ {
		total += nums[right]

		for total >= target {
			result = min(right-left+1, result)
			total -= nums[left]
			left++
		}
	}

	if result == math.MaxInt {
		return 0
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
