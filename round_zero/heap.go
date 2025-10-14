package roundzero

func adjustHeap(nums []int, idx int) {
	if idx < 0 || idx >= len(nums) {
		return
	}

	l := 2 * idx
	r := 2*idx + 1

	minIdx := idx
	if l < len(nums) && nums[l] < nums[minIdx] {
		minIdx = l
	}

	if r < len(nums) && nums[r] < nums[minIdx] {
		minIdx = r
	}

	if minIdx == idx {
		return
	}

	nums[minIdx], nums[idx] = nums[idx], nums[minIdx]

	adjustHeap(nums, minIdx)
}
