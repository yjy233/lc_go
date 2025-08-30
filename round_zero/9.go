package roundzero

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	wei := 1
	for x > 0 {
		x /= 10
		if x > 0 {
			wei *= 10
		}
	}

	x = tmp
	for x > 0 {
		if x%10 != x/wei {
			return false
		}

		x = x % wei
		x = x / 10

		wei /= 100

	}

	return true
}
