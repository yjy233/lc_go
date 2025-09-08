package roundzero

import "sort"

func nextPermutation(nums []int) {

	st := -1
	l := len(nums)
	for i := l - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			st = i - 1
			break
		}
	}

	if st == -1 {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		return
	}

	ed := st + 1
	for i := ed; i < l; i++ {
		if nums[i] > nums[st] && nums[ed] > nums[i] {
			ed = i
		}
	}

	nums[st], nums[ed] = nums[ed], nums[st]
	tmpNums := nums[st+1:]
	sort.Slice(tmpNums, func(i, j int) bool {
		return tmpNums[i] < tmpNums[j]
	})
}
