package roundzero

func longestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
		dp[i][i] = true
	}

	for i := 0; i < l-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}

	for i := 2; i < l; i++ {
		for j := 0; j+i < l; j++ {
			if s[j] == s[j+i] && dp[j+1][j+i-1] {
				dp[j][i+j] = true
			}
		}
	}

	for i := l; i >= 0; i-- {
		for j := 0; j+i < l; j++ {
			if dp[j][j+i] {
				return s[j : j+i+1]
			}
		}
	}

	return ""

}
