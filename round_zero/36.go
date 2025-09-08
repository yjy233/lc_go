package roundzero

func isValidSudoku(board [][]byte) bool {

	colM := make([]map[byte]bool, 9)
	lineM := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		colM[i] = make(map[byte]bool)
		lineM[i] = map[byte]bool{}
	}

	squareMap := make([][]map[byte]bool, 3)
	for i := 0; i < 3; i++ {
		squareMap[i] = make([]map[byte]bool, 3)
		for j := 0; j < 3; j++ {
			squareMap[i][j] = make(map[byte]bool)
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c == '.' {
				continue
			}

			if lineM[i][c] {
				return false
			}
			lineM[i][c] = true

			if colM[j][c] {
				return false
			}
			colM[j][c] = true

			sx := i / 3
			sy := j / 3

			if squareMap[sx][sy][c] {
				return false
			}
			squareMap[sx][sy][c] = true
		}
	}

	return true
}
