package roundzero

func canFinish(numCourses int, prerequisites [][]int) bool {

	res := make([]bool, numCourses)
	for i := range res {
		res[i] = true
	}

	for _, pair := range prerequisites {
		res[pair[0]] = false
	}

	st := make([]int, 0, numCourses)
	push := func(pos int) {
		st = append(st, pos)
	}
	pop := func() int {
		v := st[len(st)-1]
		st = st[:len(st)-1]
		return v
	}

	for i := range res {
		if res[i] {
			push(i)
		}
	}

	for len(st) > 0 {
		pos := pop()
		for _, pair := range prerequisites {

			src, dst := pair[0], pair[1]
			if dst != pos {
				continue
			}

			nN := true
			for _, pairV := range prerequisites {
				s, d := pairV[0], pairV[1]

				if d != src {
					continue
				}

				if !res[s] {
					nN = false
					break
				}
			}

			if nN {
				push(src)
			}
		}

	}

	for _, j := range res {
		if !j {
			return false
		}
	}

	return true
}
