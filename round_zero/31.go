package roundzero

import "sort"

func nextPermutation(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}

	ind := -1
	for i := l - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			continue
		}

		ind = i
		break
	}

	// 没毛病
	if ind == -1 {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		return
	}

	nextSt := ind + 1
	for i := nextSt; i < l; i++ {
		if nums[i] < nums[nextSt] && nums[i] > nums[ind] {
			nextSt = i
		}
	}

	nums[nextSt], nums[ind] = nums[ind], nums[nextSt]

	sort.Slice(nums[ind+1:], func(i, j int) bool {
		return nums[i] < nums[j]
	})
}
