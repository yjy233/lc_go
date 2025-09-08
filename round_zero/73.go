package roundzero

func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m <= 0 {
		return
	}

	n := len(matrix[0])
	if n <= 0 {
		return
	}

	firstLine := false
	firstCol := false

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != 0 {
				continue
			}

			matrix[i][0] = 0
			matrix[0][j] = 0

			if i == 0 {
				firstLine = true
			}
			if j == 0 {
				firstCol = true
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstLine {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if firstCol {
		for j := 0; j < m; j++ {
			matrix[j][0] = 0
		}
	}
}
