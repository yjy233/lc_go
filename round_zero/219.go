package roundzero

func containsNearbyDuplicate(nums []int, k int) bool {
	// 0 1 2 0 0
	n2C := make(map[int]int)

	for i := 0; i < min(k, len(nums)); i++ {
		if n2C[nums[i]] > 0 {
			return true
		}

		n2C[nums[i]] += 1
	}

	for i := k; i < len(nums); i++ {
		if n2C[nums[i]] > 0 {
			return true
		}

		n2C[nums[i]] += 1
		n2C[nums[i-k]] -= 1
	}

	return false
}
