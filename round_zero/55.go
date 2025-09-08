package roundzero

func canJump(nums []int) bool {

	l := len(nums)
	dp := make([]bool, l)

	if l <= 0 {
		return true
	}

	dp[0] = true

	for i := 0; i < l; i++ {
		if !dp[i] {
			continue
		}

		for j := 1; j <= nums[i]; j++ {
			if i+j >= l {
				break
			}

			dp[i+j] = true
		}
	}

	return dp[l-1]
}
