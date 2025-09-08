package roundzero

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	l := len(nums)
	if l < 3 {
		return 0
	}

	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, l-1

		tmpT := target - nums[i]
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			if k < l-1 && nums[k] == nums[k+1] {
				k--
				continue
			}

			if math.Abs(float64(target)-float64(nums[i])-float64(nums[j])-float64(nums[k])) < math.Abs(float64(target)-float64(res)) {
				res = nums[i] + nums[j] + nums[k]
			}

			if nums[j]+nums[k] > tmpT {
				k--
				continue
			}
			j++

		}

	}

	return res
}
