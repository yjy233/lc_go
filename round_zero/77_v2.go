package roundzero

func combineV2(n int, k int) [][]int {
	res := make([][]int, 0, 1<<k)
	tmp := make([][]int, 0, 1<<k)

	res = append(res, []int{})

	for i := 0; i < k; i++ {

		for _, item := range res {

			idx := len(item)
			if len(item) > 0 {
				idx = item[len(item)-1]
			}

			for j := idx + 1; j <= n; j++ {
				newItem := make([]int, len(item), len(item)+1)
				copy(newItem, item)
				newItem = append(newItem, j)
				tmp = append(tmp, newItem)
			}

		}

		tmp, res = res, tmp
		tmp = tmp[:0]
	}
	return res
}
