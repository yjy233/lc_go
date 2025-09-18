package roundzero

func maxProduct(nums []int) int {

	l := len(nums)
	maxl := make([]int, l)
	minl := make([]int, l)

	maxl[0], minl[0] = nums[0], nums[0]

	for i := 1; i < l; i++ {
		maxl[i] = max(nums[i]*maxl[i-1], nums[i]*minl[i-1], nums[i])
		minl[i] = min(nums[i]*maxl[i-1], nums[i]*minl[i-1], nums[i])
	}

	res := nums[0]
	for _, i := range maxl {
		res = max(res, i)
	}
	return res
}
