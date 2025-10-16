package roundzero

func longestSubsequenceRepeatedK(s string, k int) string {

	if len(s) < k {
		return ""
	}

	c2Cnt := make(map[byte]int)

	for i := range s {
		c2Cnt[s[i]] += 1
	}

	valid := true
	for _, cnt := range c2Cnt {
		if cnt >= k {
			continue
		}
		valid = false
		break
	}
	if valid {
		return s
	}

	idx := 0
	res := ""
	for i := range len(s) {

		if c2Cnt[s[i]] >= k {
			continue
		}

		sub := s[idx:i]
		sub = longestSubsequenceRepeatedK(sub, k)
		if len(sub) > len(res) {
			res = sub
		}
		idx = i + 1
	}

	if idx < len(s) {
		sub := s[idx:]
		sub = longestSubsequenceRepeatedK(sub, k)
		if len(sub) > len(res) {
			res = sub
		}
	}
	return res
}
