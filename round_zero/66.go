package roundzero

func plusOne(digits []int) []int {
	l := len(digits)
	if l <= 0 {
		return digits
	}

	digits[l-1] += 1
	jin := digits[l-1] / 10
	digits[l-1] = digits[l-1] % 10

	for i := l - 2; i >= 0; i-- {
		if jin == 0 {
			break
		}

		digits[i] += jin
		jin = digits[i] / 10
		digits[i] %= 10
	}

	if jin == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
