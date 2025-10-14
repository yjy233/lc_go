package roundzero

func maxSlidingWindow(nums []int, k int) []int {
	st := make([]int, 0, k+1)

	push := func(i int) {
		for len(st) > 0 {
			if nums[st[len(st)-1]] > nums[i] {
				break
			}
			st = st[:len(st)-1]
		}

		st = append(st, i)
	}

	for i := range min(k, len(nums)) {
		push(i)
	}

	res := make([]int, 0, len(nums)-k+1)
	res = append(res, nums[st[0]])

	for i := k; i < len(nums); i++ {
		push(i)

		if i-st[0] >= k {
			st = st[1:]
		}

		res = append(res, nums[st[0]])
	}

	return res
}
