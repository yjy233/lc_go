package roundzero

func numIslandsv1(grid [][]byte) int {
	res := 0
	st := make([][2]int, 0, 200)

	push := func(a, b int) {
		st = append(st, [2]int{a, b})
	}

	pop := func() (bool, [2]int) {
		if len(st) <= 0 {
			return false, [2]int{}
		}

		r := st[len(st)-1]

		st = st[:len(st)-1]
		return true, r
	}

	m := len(grid)
	if m <= 0 {
		return 0
	}
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}

			res++
			grid[i][j] = '0'
			push(i+1, j)
			push(i-1, j)
			push(i, j-1)
			push(i, j+1)

			for true {
				ok, zb := pop()
				if !ok {
					break
				}

				x := zb[0]
				y := zb[1]

				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '0' {
					continue
				}

				push(x-1, y)
				push(x+1, y)
				push(x, y+1)
				push(x, y-1)
				grid[x][y] = '0'
			}
		}
	}

	return res
}
