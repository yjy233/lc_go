package roundzero

func longestValidParentheses(s string) int {

	const NoMatch = -1
	l := len(s)
	st := make([]int, 0, l)

	pop := func() int {
		tmpL := len(st)
		if tmpL <= 0 {
			return NoMatch
		}

		v := st[tmpL-1]
		st = st[:tmpL-1]
		return v
	}

	push := func(x int) {
		st = append(st, x)
	}

	dp := make([]int, l)
	for i := range s {
		c := s[i]

		if c == '(' {
			push(i)
			continue
		}

		left := pop()
		if left == NoMatch {
			continue
		}

		dp[i] = i - left + 1
		if left > 0 {
			dp[i] += dp[left-1]
		}
	}

	res := 0
	for _, i := range dp {
		res = max(res, i)
	}
	return res
}
