package roundzero

import "sort"

func combinationSumV2(candidates []int, target int) [][]int {

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	res := make([][]int, 0, 10)
	path := make([]int, 0, len(candidates))

	helper39V1(candidates, target, 0, &path, &res)
	return res
}

func helper39V1(candidates []int, target int, idx int, path *[]int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			tmp := make([]int, len(*path))
			copy(tmp, *path)
			(*res) = append((*res), tmp)
		}
		return
	}

	if len(candidates) <= idx {
		return
	}

	for i := idx; i < len(candidates); i++ {
		if target < candidates[i] {
			return
		}

		(*path) = append((*path), candidates[i])
		helper39V1(candidates, target-candidates[i], i, path, res)
		(*path) = (*path)[:len(*path)-1]
	}
}
