package roundzero

func isInterleave(s1 string, s2 string, s3 string) bool {

	l0 := len(s1)
	l1 := len(s2)
	if l0+l1 != len(s3) {
		return false
	}

	dp := make([][]bool, l0+1)
	for i := 0; i <= l0; i++ {
		dp[i] = make([]bool, l1+1)
	}

	dp[0][0] = true
	for i := 0; i < len(s1); i++ {
		if s3[i] != s1[i] {
			break
		}
		dp[i+1][0] = true
	}
	for i := 0; i < len(s2); i++ {
		if s3[i] != s2[i] {
			break
		}
		dp[0][i+1] = true
	}

	for i := 0; i < len(s3); i++ {
		for k := 0; k <= i; k++ {
			v := i - k
			if k > l0 || v > l1 || k < 0 || v < 0 {
				continue
			}

			if dp[k][v] {
				if k < l0 && s1[k] == s3[i] {
					dp[k+1][v] = true
				}
				if v < l1 && s2[v] == s3[i] {
					dp[k][v+1] = true
				}
			}
		}
	}

	return dp[l0][l1]
}
