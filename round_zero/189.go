package roundzero

import "slices"

func Rotate189(nums []int, k int) {
	l := len(nums)
	if l == 0 {
		return
	}

	k %= l
	if k == 0 {
		return
	}

	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
