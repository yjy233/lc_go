package roundzero

/*
问题就是全搞成正的
*/
func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}

	isfu := true
	if dividend < 0 {
		isfu = divisor < 0
	} else {
		isfu = divisor > 0
		dividend = 0 - dividend
	}

	if divisor > 0 {
		divisor = 0 - divisor
	}

	tmp := []int{divisor}
	resL := []int{1}

	for dividend-divisor <= divisor {
		divisor += divisor
		tmp = append(tmp, divisor)
		resL = append(resL, resL[len(resL)-1]+resL[len(resL)-1])
	}

	res := 0
	ind := len(resL) - 1
	for ind >= 0 {
		if dividend <= tmp[ind] {
			dividend -= tmp[ind]
			res -= resL[ind]
			continue
		}

		ind--
	}

	if isfu {
		if res == -2147483648 {
			return 2147483647
		}
		return 0 - res
	}
	return res
}
