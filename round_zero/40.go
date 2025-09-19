package roundzero

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	path := make([]int, 0, len(candidates))
	res := make([][]int, 0, 100)
	use := make([]bool, len(candidates))
	helper40(candidates, 0, &path, target, &res, use)
	return res
}

func helper40(nums []int, ind int, path *[]int, target int, res *[][]int, use []bool) {
	if target == 0 {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		(*res) = append((*res), tmp)
		return
	}
	if target < 0 {
		return
	}

	for i := ind; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !use[i-1] {
			continue
		}
		if nums[i] > target {
			return
		}

		(*path) = append((*path), nums[i])
		use[i] = true
		helper40(nums, i+1, path, target-nums[i], res, use)
		use[i] = false
		(*path) = (*path)[:len(*path)-1]
	}
}
