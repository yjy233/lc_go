package roundzero

func twoSum(nums []int, target int) []int {

	nums2Ind := make(map[int]int)
	if len(nums) < 2 {
		return []int{}
	}

	nums2Ind[nums[0]] = 0

	for i := 1; i < len(nums); i++ {
		rest := target - nums[i]
		restInd, ok := nums2Ind[rest]
		if ok {
			return []int{restInd, i}
		}

		nums2Ind[nums[i]] = i
	}

	return []int{-1, -1}
}
