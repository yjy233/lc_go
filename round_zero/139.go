package roundzero

func wordBreak(s string, wordDict []string) bool {

	l := len(s)

	dp := make([]bool, l)

	for i := 0; i < l; i++ {
		if i != 0 && !dp[i-1] {
			continue
		}

		for _, word := range wordDict {
			r := i + len(word)

			if r > l || dp[r-1] {
				continue
			}

			if r <= l && s[i:r] == word {
				dp[r-1] = true
			}
		}
	}

	return dp[l-1]
}
