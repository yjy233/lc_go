package roundzero

func lengthOfLIS(nums []int) int {
	l := len(nums)
	dp := make([]int, 0, l)
	dp = append(dp, nums[0])

	for i := 1; i < l; i++ {
		tmpL := len(dp)
		if nums[i] > dp[tmpL-1] {
			dp = append(dp, nums[i])
		}

		for j := tmpL - 1; j >= 0; j-- {
			if j == len(dp)-1 {
				continue
			}

			if nums[i] > dp[j] && nums[i] < dp[j+1] {
				dp[j+1] = nums[i]
			}
		}
		dp[0] = min(dp[0], nums[i])
	}

	return len(dp)
}
