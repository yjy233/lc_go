package roundzero

import "sort"

func combinationSum(candidates []int, target int) [][]int {

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	res := make([][]int, 0)
	path := make([]int, 0, len(candidates)*5)
	helper39(&candidates, 0, &res, &path, target)
	return res
}

func helper39(nums *[]int, ind int, res *[][]int, path *[]int, target int) {
	if target == 0 {
		newTmp := make([]int, len(*path))
		copy(newTmp, *path)

		(*res) = append((*res), newTmp)
		return
	}

	if target < 0 {
		return
	}

	if ind >= len(*nums) || target < (*nums)[ind] {
		return
	}

	for i := ind; i < len(*nums); i++ {
		if target < (*nums)[ind] {
			break
		}

		(*path) = append((*path), (*nums)[i])
		helper39(nums, ind, res, path, target-(*nums)[i])
		(*path) = (*path)[:len(*nums)-1]
	}

}
