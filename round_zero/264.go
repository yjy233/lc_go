package roundzero

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	x2, x3, x5 := 1, 1, 1

	i := 2
	for i <= n {
		//fmt.Println(dp)
		t2 := dp[x2] * 2
		t3 := dp[x3] * 3
		t5 := dp[x5] * 5

		tmp := min(t2, t3, t5)
		if tmp == t2 {
			x2++
		} else if tmp == t3 {
			x3++
		} else {
			x5++
		}

		if dp[i-1] < tmp {
			dp[i] = tmp
			i++
		}
	}
	//fmt.Println(dp)
	return dp[n]
}
