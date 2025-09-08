package roundzero

func generateMatrix(n int) [][]int {

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = -1
		}
	}

	cnt := 1
	x := 0
	y := 0
	d := 0
	for cnt <= n*n {
		if res[x][y] == -1 {
			res[x][y] = cnt
			cnt++
		}

		if d%4 == 0 {
			if y < n-1 && res[x][y+1] == -1 {
				y++
			} else {
				d++
			}
			continue
		}

		if d%4 == 1 {
			if x < n-1 && res[x+1][y] == -1 {
				x++
			} else {
				d++
			}
			continue
		}

		if d%4 == 2 {
			if y > 0 && res[x][y-1] == -1 {
				y--
			} else {
				d++
			}
			continue
		}

		if x > 0 && res[x-1][y] == -1 {
			x--
		} else {
			d++
		}
	}

	return res
}
