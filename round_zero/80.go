package roundzero

func removeDuplicatesV2(nums []int) int {

	if len(nums) <= 0 {
		return 0
	}

	pre := nums[0]
	cnt := 1
	del := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == pre {
			if cnt < 2 {
				nums[i-del] = nums[i]
			} else {
				del++
			}
			cnt++
			continue
		}

		pre = nums[i]
		nums[i-del] = nums[i]
		cnt = 1
	}
	return len(nums) - del
}
