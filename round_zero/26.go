package roundzero

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	preT := nums[0]
	res := 0

	for i := 1; i < l; i++ {
		if nums[i] != preT {
			preT = nums[i]
			nums[i-res] = nums[i]
			continue
		}

		res++
	}
	return l - res
}
