package roundzero

import "math"

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}

	n := len(grid[0])
	if n <= 0 {
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			upFrom := math.MaxInt32
			lefFrom := math.MaxInt32

			if i > 0 {
				upFrom = grid[i-1][j]
			}
			if j > 0 {
				lefFrom = grid[i][j-1]
			}

			grid[i][j] += min(upFrom, lefFrom)
		}
	}
	return grid[m-1][n-1]
}
