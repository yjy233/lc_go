package roundzero

import "math"

func myAtoi(s string) int {
	s = removePreSpace(s)
	if len(s) == 0 {
		return 0
	}

	fu := true
	if s[0] == '-' {
		s = s[1:]
		fu = false
	} else if s[0] == '+' {
		s = s[1:]
	}

	res := make([]int, 0, 20)
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}

		c := int(s[i] - '0')
		if len(res) == 0 && c == 0 {
			continue
		}

		res = append(res, c)
	}

	r := 0
	overFill := false
	for i := 0; i < len(res); i++ {
		//2147483647
		if i >= 10 {
			overFill = true
			break
		}

		if i == 9 {
			if r > 214748364 {
				overFill = true
				break
			}

			if r == 214748364 {
				if !fu && res[i] > 8 {
					overFill = true
					break
				}

				if res[i] > 7 {
					overFill = true
					break
				}
			}
		}

		r *= 10
		r += res[i]
	}

	if overFill {
		if fu {
			return math.MaxInt32
		}
		return math.MinInt32
	}

	if fu {
		return r
	}
	return 0 - r
}

func removePreSpace(s string) string {
	for len(s) != 0 && s[0] == ' ' {
		s = s[1:]
	}
	return s
}
