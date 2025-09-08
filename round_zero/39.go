package roundzero

import "sort"

func combinationSum(candidates []int, target int) [][]int {

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	res := [][]int{}
	path := []int{}
	helperCombine(candidates, 0, target, &path, &res)
	return res
}

func helperCombine(nums []int, ind int, target int, path *[]int, res *[][]int) {
	//fmt.Println(ind, target, (*path))
	if target == 0 {
		newPath := make([]int, len(*path))
		copy(newPath, (*path))

		(*res) = append((*res), newPath)
		return
	}

	if ind >= len(nums) || target < 0 {
		return
	}

	for i := ind; i < len(nums); i++ {
		//fmt.Println("---",i,nums[i])
		if nums[i] > target {
			break
		}

		(*path) = append((*path), nums[i])
		helperCombine(nums, i, target-nums[i], path, res)

		(*path) = (*path)[:len(*path)-1]

	}
}
