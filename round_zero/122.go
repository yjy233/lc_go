package roundzero

func maxProfitV1(prices []int) int {

	l := len(prices)

	if l <= 1 {
		return 0
	}

	dp := make([][2]int, l)
	dp[0][1] = 0 - prices[0]

	for i := 1; i < l; i++ {

		dp[i][0] = max(dp[i-1][0], prices[i]+dp[i-1][1])

		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[l][0]
}
