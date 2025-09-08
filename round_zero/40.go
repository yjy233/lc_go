package roundzero

import "sort"

func combinationSum2(candidates []int, target int) [][]int {

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	res := &[][]int{}
	path := &[]int{}
	helperComb(&candidates, target, 0, path, res)

	return *res
}

func helperComb(candidates *[]int, target int, ind int, path *[]int, res *[][]int) {
	//fmt.Println(target)
	if target == 0 {
		newPath := make([]int, len(*path))
		copy(newPath, *path)

		(*res) = append((*res), newPath)
		return
	}

	if target < 0 {
		return
	}

	if ind >= len(*candidates) {
		return
	}

	for i := ind; i < len(*candidates); i++ {
		if (*candidates)[ind] > target {
			break
		}

		if i > ind && (*candidates)[i] == (*candidates)[i-1] {
			continue
		}

		(*path) = append((*path), (*candidates)[i])
		helperComb(candidates, target-(*candidates)[i], i+1, path, res)
		(*path) = (*path)[:len(*path)-1]
	}
}
