package roundzero

func minWindow(s string, t string) string {

	tChar2Cnt := make(map[rune]int)
	sChar2Cnt := make(map[rune]int)

	hitF := func() bool {
		for char, cnt := range tChar2Cnt {
			if sChar2Cnt[char] < cnt {
				return false
			}
		}

		return true

	}

	for _, c := range t {
		tChar2Cnt[rune(c)] += 1
	}

	res := ""

	st := 0
	ed := 1
	sChar2Cnt[rune(s[0])] = 1
	if hitF() {
		return s[:1]
	}

	for ed < len(s) {
		if !hitF() {
			sChar2Cnt[rune(s[ed])] += 1
			ed++
		}

		for hitF() {
			if res == "" || len(res) > ed-st {
				res = s[st:ed]
			}

			sChar2Cnt[rune(s[st])] -= 1
			st += 1
			continue
		}
	}

	return res
}
