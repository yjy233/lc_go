package roundzero

func maxProfitV3(prices []int) int {
	l := len(prices)

	if l <= 1 {
		return 0
	}

	// 1 持有，未交易
	// 2 交易1次不持有
	// 3 交易1次，持有1
	// 4 最终
	dp := make([][4]int, l)

	dp[0][0] = 0 - prices[0]

	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], 0-prices[i])

		if i == 1 {
			dp[i][1] = prices[i] + dp[i-1][0]
			continue
		}
		dp[i][1] = max(prices[i]+dp[i-1][0], dp[i-1][1])

		if i == 2 {
			dp[i][2] = dp[i-1][1] - prices[i]
			continue
		}
		dp[i][2] = max(dp[i-1][1]-prices[i], dp[i-1][2])

		if i == 3 {
			dp[i][3] = prices[i] + dp[i-1][2]
			continue
		}

		dp[i][3] = max(prices[i]+dp[i-1][2], dp[i-1][3])
	}

	res := max(dp[l-1][1], dp[l-1][3])
	return max(0, res)
}
