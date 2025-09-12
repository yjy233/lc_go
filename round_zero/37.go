package roundzero

func solveSudoku(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool

	var spaces [][2]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				spaces = append(spaces, [2]int{i, j})
				continue
			}

			c := board[i][j] - '1'
			line[i][c] = true
			column[j][c] = true
			block[i/3][j/3][c] = true
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos >= len(spaces) {
			return true
		}

		i, j := spaces[pos][0], spaces[pos][1]
		for c := byte(0); c < 9; c++ {
			if !line[i][c] && !column[j][c] && !block[i/3][j/3][c] {

				line[i][c] = true
				column[j][c] = true
				block[i/3][j/3][c] = true
				board[i][j] = '1' + c
				if dfs(pos + 1) {
					return true
				}

				line[i][c] = false
				column[j][c] = false
				block[i/3][j/3][c] = false
			}

		}
		return false
	}

	dfs(0)

}
