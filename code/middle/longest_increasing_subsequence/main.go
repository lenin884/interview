package longest_increasing_subsequence

//func lengthOfLIS(nums []int) int {
//	dp := make([]int, len(nums))
//
//	for i := 0; i < len(nums); i++ {
//		dp[i] = 1
//	}
//
//	for i := len(nums) - 1; i >= 0; i-- {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i] < nums[j] {
//				dp[i] = max(dp[i], 1+dp[j])
//			}
//		}
//	}
//
//	maxValue := 0
//	for _, v := range dp {
//		maxValue = max(maxValue, v)
//	}
//
//	return maxValue
//}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}
	maxValue := 0
	for _, v := range dp {
		maxValue = max(maxValue, v)
	}
	return maxValue
}
