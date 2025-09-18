package roundzero

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	res := 0
	tmp := prices[0]

	for i := 1; i < len(prices); i++ {
		res = max(res, prices[i]-tmp)

		tmp = min(prices[i], tmp)
	}

	return res
}
