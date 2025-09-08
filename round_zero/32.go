package roundzero

type IntStack []int

func (s *IntStack) Pop() int {
	if len(*s) == 0 {
		return -1
	}

	v := (*s)[len(*s)-1]

	(*s) = (*s)[:len(*s)-1]
	return v
}

func (s *IntStack) Push(c int) {
	(*s) = append((*s), c)
}

func longestValidParentheses(s string) int {

	l := len(s)

	dp := make([]int, l)

	res := 0
	st := (IntStack)(make([]int, 0, l))
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			st.Push(i)
			continue
		}

		fromInd := st.Pop()
		if fromInd == -1 {
			continue
		}

		dp[i] = i - fromInd + 1
		if fromInd > 0 {
			dp[i] += dp[fromInd-1]
		}

		res = max(res, dp[i])
	}
	return res
}
