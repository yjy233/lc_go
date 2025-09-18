package roundzero

func numDistinct(s string, t string) int {
	m := len(s)
	n := len(t)

	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 0
	if s[0] == t[0] {
		dp[0][0] = 1
	}
	for i := 1; i < m; i++ {
		dp[i][0] += dp[i-1][0]
		if s[i] == t[0] {
			dp[i][0] += 1
		}
	}

	for j := 1; j < n; j++ {
		for i := j; i < m; i++ {
			dp[i][j] += dp[i-1][j]

			if s[i] == t[j] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
