package roundzero

func getSkyline(buildings [][]int) [][]int {

	l := -1
	for i := range buildings {
		l = max(l, buildings[i][1])
	}

	if l <= 0 {
		return [][]int{}
	}

	h := make([]int, l)

	for _, t := range buildings {
		s, e, he := t[0], t[1], t[2]

		for i := s; i <= e; i++ {
			h[i] = max(h[i], he)
		}
	}

	res := make([][]int, 0, l)

	pre := 0
	for i := 0; i < l; i++ {
		if h[i] != pre {

			if h[i] > pre {
				res = append(res, []int{i, h[i]})
			} else {
				res = append(res, []int{i - 1, h[i]})
			}
			pre = h[i]
		}
	}
	return res
}
