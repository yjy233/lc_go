package roundzero

func judge(wM, tmpWM map[string]int) bool {

	for word, cnt := range wM {
		if tmpWM[word] < cnt {
			return false
		}
	}

	return true
}

func findSubstring(s string, words []string) []int {
	res := []int{}
	wCnt := len(words)
	if wCnt <= 0 {
		return res
	}

	wl := len(words[0])
	if wl <= 0 {
		return res
	}

	wM := make(map[string]int)
	for _, word := range words {
		wM[word] += 1
	}

	l := len(s)

	ind2Wm := make([]map[string]int, wl)
	for i := 0; i < wl; i++ {

		tmpWM := make(map[string]int)
		for ind := 0; ind < wCnt; ind++ {
			st := i + ind*wl
			if st+wl > l {
				break
			}

			term := s[st : st+wl]
			tmpWM[term] += 1
		}
		if judge(wM, tmpWM) {
			res = append(res, i)
		}
		ind2Wm[i] = tmpWM
	}

	for i := wl; i+wl*wCnt <= l; i++ {
		oldTerm := s[i-wl : i]
		newTerm := s[i+wl*wCnt-wl : i+wl*wCnt]

		//	_, ok := ind2Wm[i%wl][oldTerm]
		//	if !ok {
		//		panic("xxx")
		//	}

		ind2Wm[i%wl][oldTerm] -= 1
		ind2Wm[i%wl][newTerm] += 1

		if judge(wM, ind2Wm[i%wl]) {
			res = append(res, i)
		}
	}

	return res
}
