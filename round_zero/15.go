package roundzero

import (
	"sort"
)

func threeSum(nums []int) [][]int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	l := len(nums)
	res := make([][]int, 0)

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := 0 - nums[i]
		j, k := i+1, l-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			if k < l-1 && nums[k] == nums[k+1] {
				k--
				continue
			}

			if nums[j]+nums[j] > target || nums[k]+nums[k] < target {
				break
			}

			sum := nums[j] + nums[k]
			if sum == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				continue
			}

			if sum < target {
				j++
				continue
			}

			k--
		}
	}

	return res
}
