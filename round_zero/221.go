package roundzero

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	if matrix[0][0] == '1' {
		dp[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			continue
		}
	}

	for i := 1; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			continue
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dp[i][j])
		}
	}

	return ans * ans
}
