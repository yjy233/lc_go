package roundzero

func robV2(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}

	dp := make([]int, l+2)

	for i := 0; i < l; i++ {
		pre := 0
		if i >= 2 {
			pre = nums[i-2]
		}
		if i >= 3 {
			pre = max(nums[i-2], nums[i-3])
		}
		dp[i] = pre + nums[i]
	}

	tmpRes := max(nums[l-1], nums[l-2])

	for i := range l + 2 {
		dp[i] = 0
	}
	nums = append(nums, nums[:2]...)
	for i := 2; i < l+2; i++ {
		pre := 0
		if i-2 >= 2 {
			pre = nums[i-2]
		}
		if i-3 >= 2 {
			pre = max(pre, nums[i-3])
		}

		dp[i] = nums[i] + pre
	}

	return max(tmpRes, dp[l+1], dp[l])

}
