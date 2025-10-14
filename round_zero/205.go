package roundzero

func isIsomorphic(s string, t string) bool {
	sC := change205(s)
	tC := change205(t)

	if len(sC) != len(tC) {
		return false
	}

	for i := range sC {
		if sC[i] != tC[i] {
			return false
		}
	}
	return true
}

func change205(s string) []int {

	var tmp [10000]int

	res := make([]int, 0, len(s))
	cnt := 1

	for _, c := range s {
		offset := int(c)

		if tmp[offset] == 0 {
			tmp[offset] = cnt
			cnt++
		}

		res = append(res, tmp[offset])
	}

	return res
}
