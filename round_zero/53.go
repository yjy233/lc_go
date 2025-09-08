package roundzero

func maxSubArray(nums []int) int {

	l := len(nums)

	if l <= 0 {
		return 0
	}

	res := nums[0]

	for i := 1; i < l; i++ {
		nums[i] += max(0, nums[i-1])
		res = max(res, nums[i])
	}
	return res
}
