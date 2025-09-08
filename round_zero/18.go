package roundzero

import "sort"

func fourSum(nums []int, target int) [][]int {
	l := len(nums)

	res := make([][]int, 0)
	if l < 4 {
		return res
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < l-3; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		t0 := target - nums[i]

		for j := i + 1; j < l-2; j++ {
			if 3*nums[j] > t0 {
				break
			}

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			t1 := t0 - nums[j]

			x, y := j+1, l-1
			for x < y {
				if nums[x]+nums[x] > t1 || nums[y]*2 < t1 {
					break
				}

				if x > j+1 && nums[x] == nums[x-1] {
					x++
					continue
				}

				if y < l-1 && nums[y] == nums[y+1] {
					y--
					continue
				}

				if nums[x]+nums[y] == t1 {
					res = append(res, []int{
						nums[i], nums[j], nums[x], nums[y],
					})
				}

				if nums[x]+nums[y] <= t1 {
					x++
					continue
				}
				y--
			}

		}

	}

	return res
}
