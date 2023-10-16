package maximal_square

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	maxSide := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] != '0' {
				dp[i][j] = min(
					dp[i-1][j],
					dp[i][j-1],
					dp[i-1][j-1],
				) + 1
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}

	return maxSide * maxSide
}
