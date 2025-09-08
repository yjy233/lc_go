package roundzero

func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	if len(board) <= 0 || len(board[0]) <= 0 {
		return false
	}

	used := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[0]))
	}

	res := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			hlperExist(&board, &used, word, 0, &res, i, j)
		}
	}
	return res
}

func hlperExist(board *[][]byte, used *[][]bool, word string, ind int, res *bool, x, y int) {
	if *res {
		return
	}

	if (*board)[x][y] != word[ind] || (*used)[x][y] {
		return
	}

	if ind == len(word)-1 {
		*res = true
		return
	}

	(*used)[x][y] = true
	if x < len(*board)-1 {
		hlperExist(board, used, word, ind+1, res, x+1, y)
	}
	if x > 0 {
		hlperExist(board, used, word, ind+1, res, x-1, y)
	}
	if y < len((*board)[0])-1 {
		hlperExist(board, used, word, ind+1, res, x, y+1)
	}
	if y > 0 {
		hlperExist(board, used, word, ind+1, res, x, y-1)
	}

	(*used)[x][y] = false

}
