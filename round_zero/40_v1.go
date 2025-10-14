package roundzero

import "sort"

func combinationSum2V2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	path := make([]int, 0, len(candidates))
	use := make([]bool, len(candidates))
	res := make([][]int, 0, 10)

	helper40V2(candidates, &path, &use, target, 0, &res)
	return res
}

func helper40V2(nums []int, path *[]int, use *[]bool, target, idx int, res *[][]int) {
	//fmt.Println("------")
	//fmt.Println(nums, *path, *use, target, idx, *res)
	//fmt.Println(",,,,,,,,")
	if target <= 0 {
		if target == 0 {
			tmp := make([]int, len(*path))
			copy(tmp, *path)
			(*res) = append(*res, tmp)
		}
		return
	}

	if idx >= len(nums) {
		return
	}

	for i := idx; i < len(nums); i++ {
		if target < nums[i] {
			return
		}

		if i > 0 && nums[i] == nums[i-1] && !(*use)[i-1] {
			continue
		}

		(*path) = append((*path), nums[i])
		(*use)[i] = true
		helper40V2(nums, path, use, target-nums[i], i+1, res)
		(*use)[i] = false
		(*path) = (*path)[:len(*path)-1]

	}

}
