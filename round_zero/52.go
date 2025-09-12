package roundzero

func totalNQueens(n int) int {
	path := make([]int, 0, n)
	col := make([]bool, n)
	res := 0

	helper52(&path, &col, &res, n)
	return res
}

func helper52(path *[]int, col *[]bool, res *int, n int) {

	if len(*path) == n {
		(*res)++
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
		helper52(path, col, res, n)
		(*path) = (*path)[:len(*path)-1]
		(*col)[i] = false

	}

}
