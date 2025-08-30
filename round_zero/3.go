package roundzero

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l <= 1 {
		return len(s)
	}

	res := 1
	deMap := make(map[byte]bool)
	st := 0
	deMap[s[0]] = true

	i := 1
	for i < len(s) {
		c := s[i]
		has, ok := deMap[c]
		if ok && has {
			deMap[s[st]] = false
			st++
			continue
		}

		deMap[c] = true
		res = max(res, i-st+1)
		i++
	}

	return res
}
