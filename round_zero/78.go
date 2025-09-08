package roundzero

import (
	"math"
)

func subsets(nums []int) [][]int {
	l := len(nums)

	res := make([][]int, 0, int(math.Pow(2, float64(l))))
	tmp := make([][]int, 0, cap(res))

	res = append(res, []int{})

	for i := 0; i < l; i++ {

		for _, term := range res {

			tmp = append(tmp, term)

			newTerm := make([]int, len(term), len(term)+1)
			copy(newTerm, term)
			newTerm = append(newTerm, nums[i])
			tmp = append(tmp, newTerm)
		}

		res, tmp = tmp, res
		tmp = tmp[:0]
	}

	return res
}
