package roundzero

func solveNQueens(n int) [][]string {
	path := make([]int, 0, n)
	col := make([]bool, n)
	res := make([][]string, 0, n)

	helper51(&path, &col, &res, n)
	return res
}

func helper51(path *[]int, col *[]bool, res *[][]string, n int) {

	if len(*path) == n {
		frameStack := *path
		tmp := make([]string, 0, n)
		for a := 0; a < n; a++ {
			bL := make([]byte, n)
			for h := 0; h < n; h++ {
				bL[h] = '.'
			}
			bL[frameStack[a]] = 'Q'

			tmp = append(tmp, string(bL))
		}
		(*res) = append(*res, tmp)
		return
	}

	for i := 0; i < n; i++ {
		if (*col)[i] {
			continue
		}

		layer := len(*path)
		hit := true
		for j := layer - 1; j >= 0; j-- {
			if i-(layer-j) == (*path)[j] || i+(layer-j) == (*path)[j] {
				hit = false
				break
			}
		}

		if !hit {
			continue
		}

		(*col)[i] = true
		(*path) = append((*path), i)
		helper51(path, col, res, n)
		(*path) = (*path)[:len(*path)-1]
		(*col)[i] = false

	}

}
