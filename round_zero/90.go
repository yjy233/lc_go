package roundzero

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0, 1<<10)
	tmp := make([][]int, 0, 1<<10)

	res = append(res, []int{})

	for i := range nums {
		fmt.Println(res)
		for _, item := range res {
			tmp = append(tmp, item)

			if i > 0 && nums[i] == nums[i-1] {

				if len(item) == 0 {
					continue
				}

				if item[len(item)-1] != i-1 {
					continue
				}
			}

			newItem := make([]int, len(item), len(item)+1)
			copy(newItem, item)
			newItem = append(newItem, i)

			tmp = append(tmp, newItem)
		}

		res, tmp = tmp, res
		tmp = tmp[:0]
	}

	for i := range res {
		for j := range res[i] {
			res[i][j] = nums[res[i][j]]
		}
	}
	return res
}
