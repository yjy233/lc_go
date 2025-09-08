package roundzero

func search81(nums []int, target int) bool {
	return helperSearch81(nums, 0, len(nums)-1, target)
}

func helperSearch81(nums []int, l, r, target int) bool {
	if nums[l] == target || nums[r] == target {
		return true
	}

	if l >= r {
		return nums[l] == target || nums[r] == target
	}

	mid := (r + l) / 2
	if nums[mid] == target {
		return true
	}

	if mid == r || mid == l {
		return false
	}

	if nums[l] < nums[mid] {
		if binSearch81(nums, l, mid-1, target) {
			return true
		}
	} else {
		if helperSearch81(nums, l, mid-1, target) {
			return true
		}
	}

	if nums[mid] < nums[r] {
		if binSearch81(nums, mid+1, r, target) {
			return true
		}
	}

	return helperSearch81(nums, mid+1, r, target)

}

func binSearch81(nums []int, l, r, target int) bool {
	//fmt.Println("bs ", l, r, target)
	if nums[l] > target || target > nums[r] {
		return false
	}

	for l <= r {
		if nums[l] == target || nums[r] == target {
			return true
		}

		mid := (r + l) / 2
		if nums[mid] == target {
			return true
		}
		if mid == l || mid == r {
			return false
		}

		if nums[mid] > target {
			r = mid - 1
			continue
		}

		l = mid + 1

	}
	return false
}
