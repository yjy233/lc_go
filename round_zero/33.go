package roundzero

func search(nums []int, target int) int {
	return helperS(nums, target, 0, len(nums)-1)
}

func helperS(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if mid == left || mid == right {
		if nums[left] == target {
			return left
		}

		if nums[right] == target {
			return right
		}

		return -1
	}

	// 前半段有序
	if nums[left] < nums[mid] {
		bs := binSearch(nums, target, left, mid)
		if bs != -1 {
			return bs
		}

		return helperS(nums, target, mid+1, right)
	}

	bs := binSearch(nums, target, mid, right)
	if bs != -1 {
		return bs
	}
	return helperS(nums, target, left, mid-1)
}

func binSearch(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	for left <= right {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}

		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if mid == left || mid == right {
			if nums[left] == target {
				return left
			}

			if nums[right] == target {
				return right
			}

			return -1
		}

		if nums[mid] < target {
			left = mid
			continue
		}

		right = mid
	}
	return -1
}
