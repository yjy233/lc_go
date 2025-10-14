package roundzero

import "sort"

type SD struct {
	Nums []int
	Use  []bool
}

func subsetsWithDupV2(nums []int) [][]int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	l := len(nums)
	res := make([]*SD, 0, 1<<l)
	tmp := make([]*SD, 0, 1<<l)

	res = append(res, &SD{
		Nums: []int{},
		Use:  make([]bool, l),
	})

	for i := 0; i < l; i++ {

		for _, sd := range res {
			tmp = append(tmp, sd)

			if i > 0 && nums[i] == nums[i-1] && !sd.Use[i-1] {
				continue
			}

			newSD := &SD{
				Nums: make([]int, len(sd.Nums), len(sd.Nums)+1),
				Use:  make([]bool, l),
			}

			copy(newSD.Nums, sd.Nums)
			newSD.Nums = append(newSD.Nums, nums[i])
			copy(newSD.Use, sd.Use)
			newSD.Use[i] = true

			tmp = append(tmp, newSD)
		}

		tmp, res = res, tmp
		tmp = tmp[:0]
	}

	realRes := make([][]int, 0, len(res))
	for _, sd := range res {
		realRes = append(realRes, sd.Nums)
	}
	return realRes
}
