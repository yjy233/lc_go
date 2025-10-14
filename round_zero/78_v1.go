package roundzero

func subsetsV2(nums []int) [][]int {
	//sort.Slice(nums, func(i, j int) bool {
	//	return nums[i] < nums[j]
	//})

	l := len(nums)
	res := make([][]int, 0, max(10, 1<<l))
	tmp := make([][]int, 0, max(10, 1<<l))

	res = append(res, []int{})
	for i := 0; i < l; i++ {
		for _, item := range res {
			tmp = append(tmp, item)

			newItem := make([]int, len(item), len(item)+1)
			copy(newItem, item)
			newItem = append(newItem, nums[i])
			tmp = append(tmp, newItem)
		}

		tmp, res = res, tmp
		tmp = tmp[:0]
	}

	return res
}
