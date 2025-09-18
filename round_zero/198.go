package roundzero

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	for i := 2; i < len(nums); i++ {
		pre := nums[i-2]
		if i >= 3 {
			pre = max(pre, nums[i-3])
		}

		nums[i] += pre
	}
	return max(nums[len(nums)-1], nums[len(nums)-2])
}
