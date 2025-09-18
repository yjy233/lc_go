package roundzero

func numIslands(grid [][]byte) int {
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res++
				helper200(grid, i, j)
			}
		}
	}
	return res
}

func helper200(grid [][]byte, x, y int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return
	}

	if grid[x][y] == '0' {
		return
	}

	grid[x][y] = '0'

	helper200(grid, x+1, y)
	helper200(grid, x-1, y)
	helper200(grid, x, y+1)
	helper200(grid, x, y-1)
}
