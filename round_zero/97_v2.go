package roundzero

func isInterleaveV3(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)

	if l3 != l1+l2 {
		return false
	}

	dp := make([][]bool, l1+1)
	for i := range dp {
		dp[i] = make([]bool, l2+1)
	}

	dp[0][0] = true
	for i := range l1 {
		if s3[i] != s1[i] {
			break
		}
		dp[i+1][0] = true
	}

	for i := range l2 {
		if s3[i] != s2[i] {
			break
		}

		dp[0][i+1] = true
	}

	for i := range l1 {
		for j := range l2 {
			if dp[i+1][j+1] {
				continue
			}

			if !dp[i][j] {
				continue
			}

			//
			if s3[i+j] == s1[i] {
				dp[i+1][j] = true
			}

			if s3[i+j] == s2[j] {
				dp[i][j+1] = true
			}

		}
	}
	return dp[l1][l2]
}
