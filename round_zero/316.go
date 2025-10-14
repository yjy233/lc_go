package roundzero

func removeDuplicateLetters(s string) string {
	l := len(s)
	st := make([]int, 0, l)

	push := func(x int) {
		st = append(st, x)
	}

	getLast := func() int {
		if len(st) <= 0 {
			return -1
		}

		return st[len(st)-1]
	}
	pop := func() int {
		v := st[len(st)-1]
		st = st[:len(st)-1]
		return v
	}

	last := make(map[int]int)
	c2Cnt := make(map[int]int)
	cur := make(map[int]int)

	for i, c := range s {
		c2Cnt[int(c)] += 1
		last[int(c)] = i
	}

	//fmt.Println(c2Cnt)

	for i, c := range s {
		for getLast() != -1 && cur[int(c)] < 1 {
			t := getLast()
			if t < int(c) {
				break
			}

			if last[t] < i {
				break
			}

			if c2Cnt[t] > 1 {
				c2Cnt[t] -= 1
			} else {
				break
			}

			pop()
			cur[t] -= 1
		}

		if cur[int(c)] < 1 {
			push(int(c))
			cur[int(c)] += 1
		}
	}

	res := ""
	for _, c := range st {
		res += string(byte(c))
	}
	return res
}
