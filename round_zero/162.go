package roundzero

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}

	res := -1
	helper(nums, len(nums)-1, 0, &res)
	return res
}

func helper(nums []int, r, l int, res *int) {
	//fmt.Println(l,r, nums)
	if *res != -1 {
		return
	}

	mid := (r + l) / 2

	if isPeak(nums, r) {
		(*res) = r
		return
	}

	if isPeak(nums, l) {
		(*res) = l
		return
	}

	if isPeak(nums, mid) {
		(*res) = mid
		return
	}

	if l+1 >= r {
		return
	}

	if nums[mid] > nums[l] && nums[mid] < nums[r] {
		helper(nums, r, mid, res)
		return
	}

	if nums[mid] > nums[r] && nums[l] > nums[mid] {
		helper(nums, mid, l, res)
		return
	}

	helper(nums, mid, l, res)
	helper(nums, r, mid, res)
}

func isPeak(nums []int, idx int) bool {
	if idx == 0 {
		return nums[idx] > nums[1]
	}

	if idx == len(nums)-1 {
		return nums[len(nums)-1] > nums[len(nums)-2]
	}

	if idx <= 0 || idx >= len(nums)-1 {
		return false
	}

	return nums[idx] > nums[idx-1] && nums[idx] > nums[idx+1]
}
