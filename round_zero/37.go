package roundzero

func solveSudoku(board [][]byte) {
	colM := [9][9]bool{}
	lineM := [9][9]bool{}
	squareMap := [3][3][9]bool{}

	isOver := false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c == '.' {
				continue
			}

			lineM[i][c-'1'] = true
			colM[j][c-'1'] = true
			squareMap[i/3][j/3][c-'1'] = true
		}
	}

	var helperSolveSudoku func(x, y int)
	helperSolveSudoku = func(x, y int) {
		if isOver || x >= 9 || y >= 9 {
			return
		}

		if board[x][y] != '.' {
			if y < 8 {
				helperSolveSudoku(x, y+1)
				return
			} else {
				helperSolveSudoku(x+1, 0)
			}
			return
		}

		for c := (byte)('1'); c <= '9'; c++ {
			if colM[y][c-'1'] || lineM[x][c-'1'] || squareMap[x/3][y/3][c-'1'] {
				continue
			}

			board[x][y] = c
			if x == 8 && y == 8 {
				isOver = true
				return
			}

			lineM[x][c-'1'] = true
			colM[y][c-'1'] = true
			squareMap[x/3][y/3][c-'1'] = true

			if y < 8 {
				helperSolveSudoku(x, y+1)
			} else {
				helperSolveSudoku(x+1, 0)
			}
			if isOver {
				return
			}

			board[x][y] = '.'
			lineM[x][c-'1'] = false
			colM[y][c-'1'] = false
			squareMap[x/3][y/3][c-'1'] = false
		}
	}

	helperSolveSudoku(0, 0)
	//helperSolveSudoku(&board, &isOver, &colM, &lineM, &squareMap, 0, 0)
}
