package roundzero

import "sort"

func permute(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	path := make([]int, 0, len(nums))
	res := make([][]int, 0, 100)
	use := make([]bool, len(nums))
	helper46(nums, use, &path, &res)
	return res
}

func helper46(nums []int, use []bool, path *[]int, res *[][]int) {
	if len(*path) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, *path)
		(*res) = append((*res), tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if use[i] || (i > 0 && nums[i] == nums[i-1] && !use[i-1]) {
			continue
		}

		use[i] = true
		(*path) = append((*path), nums[i])

		helper46(nums, use, path, res)
		use[i] = false
		(*path) = (*path)[:len(*path)-1]
	}

}
