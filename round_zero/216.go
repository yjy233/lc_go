package roundzero

func combinationSum3(k int, n int) [][]int {
	if n > 45 || k > n {
		return [][]int{}
	}

	res := make([][]int, 0)

	path := make([]int, 0, 10)

	var dfs func(int, int)
	dfs = func(pos, taget int) {
		if len(path) > k {
			return
		}
		if taget == 0 && len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		if pos >= 10 {
			return
		}

		for i := pos; i <= 9; i++ {
			path = append(path, i)
			dfs(i+1, taget-i)
			path = path[:len(path)-1]
		}
	}
	dfs(1, n)
	return res
}
