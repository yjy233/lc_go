package roundzero

import (
	"math"
	"sort"
)

func permuteUnique(nums []int) [][]int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	l := len(nums)
	resN := math.Pow(2.0, float64(l))
	res := make([][]int, 0, int(resN))
	use := make([]bool, l)
	path := make([]int, 0, l)

	helperPermuteUnique(&res, &use, &path, &nums)
	return res
}

func helperPermuteUnique(res *[][]int, use *[]bool, path *[]int, nums *[]int) {
	//fmt.Println("mmmmmm", (*path))
	if len(*path) == len(*nums) {
		//fmt.Println("0-0000000")
		tmpPath := make([]int, len(*path))
		//fmt.Println(tmpPath)
		copy(tmpPath, *path)
		*res = append(*res, tmpPath)
		return
	}

	for i := 0; i < len(*nums); i++ {
		// 重复问题，上一个如果没用直接跳过
		if (*use)[i] || (i > 0 && (*nums)[i-1] == (*nums)[i] && !(*use)[i-1]) {
			continue
		}

		(*use)[i] = true
		(*path) = append((*path), (*nums)[i])
		helperPermuteUnique(res, use, path, nums)
		(*path) = (*path)[:len(*path)-1]
		(*use)[i] = false
	}

}
