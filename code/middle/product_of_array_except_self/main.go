package product_of_array_except_self

// productExceptSelf time complexity O(n)
// space complexity O(1) - if we don't count the result array
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}
