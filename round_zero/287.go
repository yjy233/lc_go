package roundzero

func findDuplicate(nums []int) int {

	i := 0
	n := len(nums) - 1
	for i < n+1 {
		if nums[i] == i+1 {
			i++
			continue
		}

		tmp := nums[i]
		if nums[tmp-1] == tmp {
			return tmp
		}

		nums[i], nums[tmp-1] = nums[tmp-1], nums[i]

	}
	return n
}
