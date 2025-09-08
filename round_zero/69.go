package roundzero

// x*x + 2x + 1
func mySqrt(x int) int {

	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	if x < 4 {
		return 1
	}

	i := 2
	x2f := 4
	for x >= i {
		if x-x2f < 2*i+1 {
			return i
		}

		x2f = (i + 1) * (i + 1)

		i++
	}

	return i
}
