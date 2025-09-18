package roundzero

func maxProfit155(k int, prices []int) int {
	if k <= 0 {
		return 0
	}

	l := len(prices)

	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, 2*k)
	}

	dp[0][0] = 0 - prices[0]

	for i := 1; i < l; i++ {
		for j := 0; j < min(2*k, 2*i+1); j++ {
			if j%2 == 0 {
				pre := 0
				if j > 0 {
					pre = dp[i-1][j-1]
				}
				dp[i][j] = max(dp[i-1][j], pre-prices[i])

				// 爬升
				if j == 2*i {
					dp[i][j] = pre - prices[i]
				}
			} else {

				if j == 2*i-1 {
					dp[i][j] = prices[i] + dp[i-1][j-1]
				} else {
					dp[i][j] = max(dp[i-1][j], prices[i]+dp[i-1][j-1])
				}
			}
		}
	}

	res := 0
	for _, i := range dp[l-1] {
		res = max(res, i)
	}
	return res
}
