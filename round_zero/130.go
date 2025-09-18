package roundzero

func solve(board [][]byte) {

	m := len(board)
	if m <= 0 {
		return
	}
	n := len(board[0])
	if n <= 0 {
		return
	}

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			helper130(&board, 'O', 'L', i, 0)
		}
		if board[i][n-1] == 'O' {
			helper130(&board, 'O', 'L', i, n-1)
		}
	}
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			helper130(&board, 'O', 'L', 0, i)
		}
		if board[m-1][i] == 'O' {
			helper130(&board, 'O', 'L', m-1, i)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'L' {
				board[i][j] = 'O'
			}
		}
	}
}

func helper130(board *[][]byte, src, dst byte, i, j int) {

	if (*board)[i][j] != src {
		return
	}

	(*board)[i][j] = dst

	if i < len(*board)-1 {
		helper130(board, src, dst, i+1, j)
	}
	if i > 0 {
		helper130(board, src, dst, i-1, j)
	}

	if j < len((*board)[0])-1 {
		helper130(board, src, dst, i, j+1)
	}

	if j > 0 {
		helper130(board, src, dst, i, j-1)
	}

}
