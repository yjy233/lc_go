package roundzero

func minSubArrayLen(target int, nums []int) int {
	res := -1

	setRes := func(r int) {
		if res == -1 {
			res = r
			return
		}

		res = min(res, r)
	}

	sum := int64(0)
	for i := range nums {
		sum += int64(nums[i])
	}

	if sum < int64(target) {
		return 0
	}

	res = len(nums)

	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] <= nums[r] {
			if sum-int64(nums[l]) < int64(target) {
				break
			}
			l++
		} else {
			if sum-int64(nums[r]) < int64(target) {
				break
			}
			r--
		}

		setRes(r - l + 1)
	}
	return res
}
