package triangle

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	dp := make([]int, len(triangle))

	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if i == len(triangle)-1 {
				dp[j] = triangle[i][j]
			} else {
				dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
			}
		}
	}

	return dp[0]
}
