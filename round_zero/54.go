package roundzero

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m <= 0 {
		return []int{}
	}

	n := len(matrix[0])
	if n <= 0 {
		return []int{}
	}

	res := make([]int, 0, m*n)

	x := 0
	y := 0
	d := 0
	usedFlag := -23423

	for len(res) < m*n {
		if matrix[x][y] != usedFlag {
			res = append(res, matrix[x][y])
			matrix[x][y] = usedFlag
		}

		if d%4 == 0 {
			if y < n-1 && matrix[x][y+1] != usedFlag {
				y++
				continue
			}
			d++
			continue
		}

		if d%4 == 1 {
			if x < m-1 && matrix[x+1][y] != usedFlag {
				x++
				continue
			}
			d++
			continue
		}

		if d%4 == 2 {
			if y > 0 && matrix[x][y-1] != usedFlag {
				y--
				continue
			}
			d++
			continue
		}

		if x > 0 && matrix[x-1][y] != usedFlag {
			x--
			continue
		}
		d++
	}

	return res
}
