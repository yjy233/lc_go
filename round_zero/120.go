package roundzero

func minimumTotal(triangle [][]int) int {

	for i := 1; i < len(triangle); i++ {

		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][0]
				continue
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][len(triangle[i-1])-1]
				continue
			}
			triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
		}
	}

	l := len(triangle)
	le := len(triangle[l-1])
	res := triangle[l-1][0]
	for i := 1; i < le; i++ {
		res = min(res, triangle[l-1][i])
	}
	return res
}
