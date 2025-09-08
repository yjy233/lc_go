package roundzero

func combine(n int, k int) [][]int {
	res := make([][]int, 0, k)
	nums := make([]int, 0, k)

	helperCombine77(&nums, 1, &res, k, n)
	return res
}

func helperCombine77(nums *[]int, ind int, res *[][]int, k int, n int) {

	if len(*nums) == k {
		tmp := make([]int, k)
		copy(tmp, *nums)
		(*res) = append((*res), tmp)
		return
	}

	if ind > n {
		return
	}

	helperCombine77(nums, ind+1, res, k, n)

	(*nums) = append((*nums), ind)
	helperCombine77(nums, ind+1, res, k, n)
	(*nums) = (*nums)[0 : len(*nums)-1]
}
