package roundzero

func firstMissingPositiveV3(nums []int) int {
	length := len(nums)

	for i := 0; i < len(nums); i++ {
		t := i + 1

		for nums[i] != t {
			tmp := nums[i]
			if tmp <= 0 || tmp > length {
				break
			}

			if nums[tmp-1] == tmp {
				break
			}

			nums[i], nums[tmp-1] = nums[tmp-1], nums[i]
		}

	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return length + 1
}
