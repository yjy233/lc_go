package roundzero

func singleNumber(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}

	res := nums[0]
	for i := 1; i < l; i++ {
		res ^= nums[i]
	}
	return res
}
