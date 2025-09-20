package roundzero

func isHappy(n int) bool {
	r := map[int]bool{}

	for !r[n] {
		res := 0
		r[n] = true
		for n > 0 {
			w := n % 10
			res += w * w
			n /= 10
		}

		if res == 1 {
			return true
		}
		n = res
	}
	return false
}
