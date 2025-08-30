package roundzero

func reverse(x int) int {
	isF := true
	if x == -2147483648 {
		return 0
	}

	if x < 0 {
		isF = false
		x = 0 - x
	}

	wcnt := 0
	res := 0
	for x > 0 {
		tmp := x % 10
		if wcnt >= 8 {
			if res > 214748364 {
				return 0
			}

			if res == 214748364 {
				if (isF && tmp > 7) || (!isF && tmp > 8) {
					return 0
				}

			}

		}

		x = x / 10
		res *= 10
		res += tmp
		wcnt += 1
	}

	if isF {
		return res
	}
	return 0 - res
}
