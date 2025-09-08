package roundzero

func jump(nums []int) int {
	l := len(nums)
	if l <= 0 {
		return 0
	}

	dp := make([]int, l)
	for i := 1; i < l; i++ {
		dp[i] = -1
	}

	//res := -1
	for i := 0; i < l; i++ {
		if dp[i] == -1 {
			continue
		}

		for j := 1; j <= nums[i]; j++ {
			ind := i + j
			if ind >= l {
				continue
				//if res == -1 || res < dp[i]+1 {
				//	res = dp[i] + 1
				//}
				//continue
			}

			if dp[ind] == -1 || dp[ind] > dp[i]+1 {
				dp[ind] = dp[i] + 1
			}
		}
	}

	return dp[l-1]
}
