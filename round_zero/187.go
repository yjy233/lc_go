package roundzero

func findRepeatedDnaSequences(s string) []string {
	s2Cnt := make(map[string]int)

	for i := 0; i+10 <= len(s); i++ {
		subStr := s[i : i+10]
		s2Cnt[subStr] += 1
	}

	res := make([]string, 0, len(s2Cnt))
	for subStr, cnt := range s2Cnt {
		if cnt >= 2 {
			res = append(res, subStr)
		}
	}
	return res
}
