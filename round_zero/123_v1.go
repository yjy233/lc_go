package roundzero

func maxProfitV5(prices []int) int {

	l := len(prices)

	if l <= 1 {
		return 0
	}

	dp := make([][4]int, l)
	dp[0][0] = 0 - prices[0]

	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], 0-prices[i])
		if i == 1 {
			dp[i][1] = prices[i] + dp[i-1][0]
			continue
		}
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])

		if i == 2 {
			dp[i][2] = dp[i-1][1] - prices[i]
			continue
		}
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])

		if i == 3 {
			dp[i][3] = dp[i][2] + prices[i]
			continue
		}
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}

	//for i := range dp {
	//   fmt.Println(dp[i])
	//}

	return max(dp[l-1][3], dp[l-1][1], 0)
}
