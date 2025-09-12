package roundzero

func isScramble(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	n := len(s1)
	if n <= 1 {
		return s1 == s2
	}

	dp := make([][][]bool, n)
	for i := range dp {
		dp[i] = make([][]bool, n)
		for j := range dp[i] {
			dp[i][j] = make([]bool, n+1)
		}
	}

	for i := range n {
		for j := range n {
			if s1[i] == s2[j] {
				dp[i][j][1] = true
			}
		}
	}

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			for j := 0; j <= n-l; j++ {
				for k := 1; k < l; k++ {
					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}

					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}

				}
			}
		}

	}

	return dp[0][0][n]
}
