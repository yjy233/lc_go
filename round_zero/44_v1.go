package roundzero

func isMatchV1(s string, p string) bool {
	l1 := len(s)
	l2 := len(p)

	dp := make([][]bool, l1+1)
	for i := range dp {
		dp[i] = make([]bool, l2)
	}

	dp[0][0] = true

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if dp[i+1][j+1] {
				continue
			}

			if p[j] == '?' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
				continue
			}

			if p[j] != '*' {
				continue
			}

			if j == 0 {
				for k := 0; k <= l1; k++ {
					dp[k][j+1] = true
				}
				continue
			}

			if dp[i][j] {
				dp[i][j+1] = true

				for k := i; k <= l1; k++ {
					dp[k][j+1] = true
				}
			}
		}
	}
	return dp[l1][l2]
}
