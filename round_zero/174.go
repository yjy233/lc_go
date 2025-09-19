package roundzero

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m <= 0 {
		return 0
	}
	n := len(dungeon[0])
	if n <= 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[m-1][n-1] = max(0, 0-dungeon[m-1][n-1])

	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = max(0, 0-dungeon[m-1][i]+dp[m-1][i+1])
	}

	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(0, 0-dungeon[i][n-1]+dp[i+1][n-1])
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = max(0, 0-dungeon[i][j]+min(dp[i+1][j], dp[i][j+1]))
		}
	}

	return dp[0][0] + 1
}
