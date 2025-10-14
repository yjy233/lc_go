package roundzero

func maxProfit309(prices []int) int {

	l := len(prices)
	if l <= 1 {
		return 0
	}

	// 0 - 可交易
	// 1 - 持有1
	// 2 - 不可交易

	dp := make([][3]int, l)
	dp[0][1] = 0 - prices[0]

	for i := 1; i < l; i++ {
		dp[i][0] = dp[i-1][0]

		if i >= 2 {
			dp[i][0] = max(dp[i][0], dp[i-1][2])
		}

		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])

		dp[i][2] = dp[i-1][1] + prices[i]
	}

	//fmt.Println(dp)

	return max(dp[l-1][0], dp[l-1][2])
}
