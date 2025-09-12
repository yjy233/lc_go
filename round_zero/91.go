package roundzero

func numDecodings(s string) int {
	l := len(s)

	if l <= 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, l)
	dp[0] = 1

	for i := 1; i < l; i++ {
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {

				return 0
			}

			if i == 1 {
				dp[i] = 1
			} else {
				dp[i] += dp[i-2]
			}

			continue
		}

		dp[i] += dp[i-1]
		need := s[i-1] == '1'
		if s[i-1] == '2' && s[i] <= '6' {
			need = true
		}

		if need {
			if i == 1 {
				dp[i] += 1
			} else {
				dp[i] += dp[i-2]
			}
		}

	}
	return dp[l-1]
}
