package roundzero

type o301V1 struct {
	cnt int
	raw []byte
}

func copyO301V1(r *o301V1, i int) *o301V1 {
	rr := &o301V1{
		cnt: r.cnt,
		raw: make([]byte, len(r.raw), i),
	}

	copy(rr.raw, r.raw)

	return rr
}

func removeInvalidParentheses301(s string) []string {
	rL := 0
	tmp := make([]string, 0, 100)

	var dfs func(*o301V1, int)

	dfs = func(r *o301V1, idx int) {
		if idx >= len(s) {
			if len(r.raw) > rL && r.cnt == 0 {
				tmp = tmp[:0]
				tmp = append(tmp, string(r.raw))
				rL = len(r.raw)
			} else if len(r.raw) == rL && r.cnt == 0 {
				tmp = append(tmp, string(r.raw))
			}
			return
		}

		// 1 2 3 4
		// 1  2
		if len(r.raw)+len(s)-idx < rL {
			return
		}

		if s[idx] == '(' || s[idx] == ')' {
			nr := copyO301V1(r, len(s))
			dfs(nr, idx+1)
		}

		switch s[idx] {
		case '(':
			r.cnt++
			r.raw = append(r.raw, s[idx])
			dfs(r, idx+1)
		case ')':
			if r.cnt > 0 {
				r.cnt--
				r.raw = append(r.raw, s[idx])
				dfs(r, idx+1)
			}
		default:
			r.raw = append(r.raw, s[idx])
			dfs(r, idx+1)
		}

	}

	tmpr := &o301V1{
		cnt: 0,
		raw: make([]byte, 0, len(s)),
	}
	dfs(tmpr, 0)

	m2C := make(map[string]bool)
	realR := make([]string, 0, len(tmp))

	for _, i := range tmp {
		if m2C[i] {
			continue
		}

		m2C[i] = true
		realR = append(realR, i)
	}

	return realR
}
