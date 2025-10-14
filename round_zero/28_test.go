package roundzero

import "testing"

func TestStr(t *testing.T) {
	needle := "abab"
	haystack := "ababab"
	dp := buildNeedle(needle)

	res := make([]int, 0)

	i := 0
	j := 0

	for i < len(haystack) {
		//fmt.Println(i,j)
		if needle[j] == haystack[i] {
			// abca
			// bca

			if j == len(needle)-1 {
				res = append(res, i-len(needle)+1)
				j = dp[j]
				continue
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

	t.Log(res)
}
