package roundzero

func hammingWeight(n int) int {
	res := 0

	for n > 0 {
		if n%2 == 1 {
			res++
		}

		n = n >> 1
	}
	return res
}
