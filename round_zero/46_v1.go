package roundzero

import "sort"

func permuteV1(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	use := make([]bool, len(nums))
	path := make([]int, 0, len(nums))
	res := make([][]int, 0, 1<<len(nums))

	dfs46(nums, &use, &path, &res)
	return res
}

func dfs46(nums []int, use *[]bool, path *[]int, res *[][]int) {
	if len(*path) == len(nums) {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		(*res) = append((*res), tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*use)[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !(*use)[i-1] {
			continue
		}

		(*path) = append((*path), nums[i])
		(*use)[i] = true
		dfs46(nums, use, path, res)
		(*use)[i] = false
		(*path) = (*path)[:len(*path)-1]
	}

}
