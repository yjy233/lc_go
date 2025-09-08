package roundzero

func removeElement(nums []int, val int) int {

	cnt := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == val {
			cnt++
			continue
		}

		nums[i-cnt] = nums[i]
	}

	return l - cnt
}
