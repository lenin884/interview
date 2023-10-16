package house_robber

func rob(nums []int) int {
	rob1, rob2 := 0, 0

	// rob1 is the max value of the previous house
	for i := 0; i < len(nums); i++ {
		// rob1 is the max value of the previous house
		current := max(rob1+nums[i], rob2)
		// rob2 is the max value of the current house
		rob1 = rob2
		// current is the max value of the current house
		rob2 = current
	}

	return rob2
}
