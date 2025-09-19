package roundzero

func combine(n int, k int) [][]int {
	use := make([]bool, n+1)

	path := make([]int, 0, k)

	res := make([][]int, 0, 256)
	helper77(&path, use, 1, &res, n, k)
	return res

}

func helper77(path *[]int, use []bool, ind int, res *[][]int, n, k int) {
	if len(*path) == k {
		tmp := make([]int, k)
		copy(tmp, *path)
		(*res) = append((*res), tmp)
		return
	}

	for i := ind; i <= n; i++ {
		if use[i] {
			continue
		}

		(*path) = append((*path), i)
		use[i] = true
		helper77(path, use, i+1, res, n, k)
		use[i] = false
		(*path) = (*path)[:len(*path)-1]
	}
}
