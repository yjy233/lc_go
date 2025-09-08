package roundzero

func minDistance(word1 string, word2 string) int {

	l0 := len(word1)
	l1 := len(word2)
	dp := make([][]int, l0+1)
	for i := range l0 + 1 {
		dp[i] = make([]int, l1+1)
	}

	dp[0][0] = 0
	for i := 1; i <= l0; i++ {
		dp[i][0] = i
	}

	for i := 1; i <= l1; i++ {
		dp[0][i] = i
	}

	for i := 0; i < l0; i++ {
		for j := 0; j < l1; j++ {

			tmp := -1
			if word1[i] == word2[j] {
				tmp = dp[i][j]
			} else {
				tmp = dp[i][j] + 1
			}

			tmp = min(tmp, dp[i+1][j]+1)
			tmp = min(tmp, dp[i][j+1]+1)

			dp[i+1][j+1] = tmp
		}
	}
	return dp[l0][l1]
}
