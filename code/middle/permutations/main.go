package permutations

func permute(nums []int) [][]int {
	var result [][]int

	backtrack(nums, &result, 0)

	return result
}

func backtrack(nums []int, result *[][]int, idx int) {
	if idx == len(nums) {
		*result = append(*result, append([]int{}, nums...))
		return
	}

	for i := idx; i < len(nums); i++ {
		nums[i], nums[idx] = nums[idx], nums[i]
		backtrack(nums, result, idx+1)
		nums[i], nums[idx] = nums[idx], nums[i]
	}

	return
}
