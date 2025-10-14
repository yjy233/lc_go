package roundzero

func minSubArrayLen(target int, nums []int) int {

	s := 0
	for _, i := range nums {
		s += i
	}

	if s < target {
		return -1
	}
	res := len(nums)

	l := 0
	r := 0
	s = nums[0]
	if s >= target {
		return 1
	}

	for r < len(nums) {
		if l <= r && s >= target {
			res = min(res, r-l+1)

			s -= nums[l]
			l++
			continue
		}

		r++
		if r >= len(nums) {
			break
		}
		s += nums[r]
		if s >= target {
			res = min(res, r-l+1)
		}
	}
	return res
}
