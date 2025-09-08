package roundzero

import "math"

func permute(nums []int) [][]int {
	l := len(nums)
	resN := math.Pow(2.0, float64(l))
	res := make([][]int, 0, int(resN))
	use := make([]bool, l)
	path := make([]int, 0, l)

	helperPermute(&res, &use, &path, &nums)
	return res
}

func helperPermute(res *[][]int, use *[]bool, path *[]int, nums *[]int) {

	if len(*path) == len(*nums) {
		//fmt.Println("0-0000000")
		tmpPath := make([]int, len(*path))
		//fmt.Println(tmpPath)
		copy(tmpPath, *path)
		*res = append(*res, tmpPath)
		return
	}

	for i := 0; i < len(*nums); i++ {
		if (*use)[i] {
			continue
		}

		//fmt.Println("ind ",i)

		(*use)[i] = true
		(*path) = append((*path), (*nums)[i])
		helperPermute(res, use, path, nums)
		(*path) = (*path)[:len(*path)-1]
		(*use)[i] = false
	}
}
