package jump_game

func canJump(nums []int) bool {
	// edge case
	if len(nums) == 1 {
		return true
	}

	// max index we can reach
	maxIndex := 0

	// [0, 3, 2, 1, 0, 1]
	for i := 0; i < len(nums); i++ {
		// we can't reach current index
		if i > maxIndex {
			return false
		}

		// update max index
		if i+nums[i] > maxIndex {
			maxIndex = i + nums[i]
		}
	}

	return true
}
