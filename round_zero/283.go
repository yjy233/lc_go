package roundzero

func moveZeroes(nums []int) {

	tmp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			tmp++
			continue
		}

		nums[i-tmp] = nums[i]
		if i-tmp != i {
			nums[i] = 0

		}
	}

}
