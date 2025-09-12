package roundzero

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m <= 0 {
		return 0
	}
	n := len(matrix[0])
	if n <= 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		pre := 0
		for j := 0; j < m; j++ {
			if matrix[j][i] == '1' {
				pre++
			} else {
				pre = 0
			}

			dp[j][i] = pre
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		res = max(res, largestRectangleArea(dp[i]))
	}
	return res
}
