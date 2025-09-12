package roundzero

func numTrees(n int) int {

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range n + 1 {
		dp[i][i] = 1
	}

	helper96(1, n, &dp)
	return dp[1][n]
}

func helper96(st, ed int, dp *[][]int) int {
	if st >= ed {
		return 1
	}

	if (*dp)[st][ed] != 0 {
		return (*dp)[st][ed]
	}

	res := 0
	for i := st; i <= ed; i++ {
		lc := helper96(st, i-1, dp)
		rc := helper96(i+1, ed, dp)

		res += (lc * rc)
	}
	(*dp)[st][ed] = res
	return res
}
