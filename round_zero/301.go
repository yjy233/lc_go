package roundzero

type o301 struct {
	cnt  int
	rond int
	raw  []byte
}

func copyO301(r *o301, i int) *o301 {
	rr := &o301{
		cnt:  r.cnt,
		rond: r.rond,
		raw:  make([]byte, len(r.raw), i),
	}

	copy(rr.raw, r.raw)

	return rr
}

func removeInvalidParentheses(s string) []string {
	res := make([]*o301, 0, 512)
	tmp := make([]*o301, 0, 512)
	res = append(res, &o301{
		raw: make([]byte, 0),
	})

	leftC := make([]int, len(s))
	rightC := make([]int, len(s))
	rC := 0
	lC := 0
	for i := len(s) - 1; i >= 0; i-- {
		rightC[i] = rC
		if s[i] == ')' {
			rC++
		}

		leftC[i] = lC
		if s[i] == '(' {
			lC++
		}
	}

	for i, c := range s {

		for _, r := range res {
			if c == '(' {
				if r.cnt+1+leftC[i] >= rightC[i] {

					newR := copyO301(r, len(s))
					newR.rond = i
					newR.cnt++
					newR.raw = append(newR.raw, byte(c))
					tmp = append(tmp, newR)
				}
			} else if c == ')' {
				if r.cnt > 0 {
					newR := copyO301(r, len(s))
					newR.rond = i
					newR.cnt--
					newR.raw = append(newR.raw, byte(c))
					tmp = append(tmp, newR)
				}
			} else {
				r.rond++
				r.raw = append(r.raw, byte(c))
			}

			tmp = append(tmp, r)
		}

		res, tmp = tmp, res
		tmp = tmp[:0]

		m2C := make(map[int]*o301)
		for _, item := range res {
			if m2C[item.cnt] == nil {
				m2C[item.cnt] = item
				continue
			}

			if len(item.raw) > len(m2C[item.cnt].raw) {
				m2C[item.cnt] = item
			}
		}

		res = res[:0]
		for _, item := range m2C {
			res = append(res, item)
		}

	}

	mCnt := 0
	for _, item := range res {
		if item.cnt == 0 && len(item.raw) > mCnt {
			mCnt = len(item.raw)
		}
	}

	sCnt := make(map[string]bool)
	rs := make([]string, 0, 1<<len(s))
	for _, item := range res {
		if item.cnt == 0 && len(item.raw) == mCnt && !sCnt[string(item.raw)] {

			rs = append(rs, string(item.raw))
			sCnt[string(item.raw)] = true
		}
	}
	return rs
}
