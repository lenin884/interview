package jump_game_2

// jump time complexity O(n)
// space complexity O(1)
func jump(nums []int) int {
	// edge case
	if len(nums) == 1 {
		return 0
	}

	jumps := 0      // number of jumps needed to reach the end
	currentEnd := 0 // current end of the jump
	farthest := 0   // farthest index we can reach

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == currentEnd {
			jumps++
			currentEnd = farthest

			// if we can reach the end of the array with current jump
			if currentEnd >= len(nums)-1 {
				return jumps
			}
		}
	}

	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
