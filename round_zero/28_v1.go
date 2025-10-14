package roundzero

func strStrV1(haystack string, needle string) int {
	dp := buildNeedle(needle)

	i := 0
	j := 0

	for i < len(haystack) {
		//fmt.Println(i,j)
		if needle[j] == haystack[i] {
			// abca
			// bca

			if j == len(needle)-1 {
				return i - len(needle) + 1
			}

			i++
			j++
			continue
		}

		if j == 0 {
			i++
			continue
		}

		j = dp[j]
	}

	return -1
}

func buildNeedle(needle string) []int {
	l := len(needle)

	res := make([]int, l)

	i := 1
	j := 0
	for i < l {
		res[i] = j
		if needle[i] == needle[j] {
			i++
			j++
			continue
		}

		for j > 0 && needle[j] != needle[i] {
			j = res[j]
		}

		if j > 0 || needle[j] == needle[i] {
			j++
		}
		i++
	}
	return res
}
