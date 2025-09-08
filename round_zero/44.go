package roundzero

func isMatchV2(s string, p string) bool {

	ls := len(s)
	lp := len(p)

	dp := make([][]bool, ls+1)
	for i := 0; i < ls+1; i++ {
		dp[i] = make([]bool, lp+1)
	}

	dp[0][0] = true

	for i := 0; i < ls; i++ {
		for j := 0; j < lp; j++ {
			if p[j] != '?' && p[j] != '*' {
				if dp[i][j] && p[j] == s[i] {
					dp[i+1][j+1] = true
				}
				continue
			}

			if p[j] == '?' {
				dp[i+1][j+1] = dp[i][j]
				continue
			}

			for k := 0; i <= ls; k++ {
				if dp[k][i] {
					for c := k; c <= ls; c++ {
						dp[c][i+1] = true
					}
				}
			}
		}
	}

	return dp[ls][lp]
}
