package roundzero

func lengthOfLongestSubstringV1(s string) int {

	l := len(s)
	if l <= 1 {
		return l
	}

	res := 1
	ed := 0
	st := 0

	m := map[byte]int{
		s[0]: 1,
	}

	//judge := func() bool {
	//	for _, v := range m {
	//		if v >= 2 {
	//			return false
	//		}
	//	}
	//	return true
	//}

	for ed < l {
		ed++
		if ed >= l {
			break
		}

		c := s[ed]
		for m[c] >= 1 {
			m[s[st]] -= 1
			st++
		}

		m[c] += 1

		res = max(res, ed-st+1)
	}

	return res
}
