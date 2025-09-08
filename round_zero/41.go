package roundzero

func firstMissingPositive(nums []int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == i+1 {
			continue
		}

		if nums[i] > len(nums) {
			continue
		}

		for nums[i] != i+1 {
			tmp := nums[i]
			if tmp <= 0 || tmp > len(nums) {
				break
			}

			if nums[tmp-1] == tmp {
				break
			}
			nums[i] = nums[tmp-1]
			nums[tmp-1] = tmp
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
