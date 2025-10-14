package roundzero

func minCut(s string) int {
	l := len(s)

	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
	}

	for i := range l {
		dp[i][i] = 1
	}

	for i := 1; i < l; i++ {
		if s[i-1] == s[i] {
			dp[i-1][i] = 1
		}
	}

	for w := 3; w <= l; w++ {
		for i := 0; i+w-1 < l; i++ {
			if s[i] == s[i+w-1] {
				dp[i][i+w-1] = min(dp[i+1][i+w-1]+1, dp[i][i+w-2]+1, dp[i+1][i+w-2])
				continue
			}
			dp[i][i+w-1] = min(dp[i+1][i+w-1]+1, dp[i][i+w-2]+1)
		}
	}

	return dp[0][l-1]
}
