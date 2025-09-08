package roundzero

func buildNextInd(needle string) []int {
	l := len(needle)
	dp := make([]int, l)
	ind := 0
	for i := 1; i < l; i++ {
		dp[i] = ind

		if needle[ind] == needle[i] {
			ind++
			continue
		}

		for ind > 0 && needle[ind] != needle[i] {
			ind = dp[ind]
		}

		if needle[ind] == needle[i] {
			ind++
		}
	}

	return dp
}

func strStr(haystack string, needle string) int {
	dp := buildNextInd(needle)

	i := 0
	j := 0

	for j < len(haystack) {
		if haystack[j] == needle[i] {
			i++
			if i >= len(needle) {
				return j - len(needle) + 1
			}
			j++
			continue
		}

		if i == 0 {
			j++
			continue
		}

		i = dp[i]

	}
	return -1
}
