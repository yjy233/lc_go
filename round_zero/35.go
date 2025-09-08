package roundzero

func searchInsert(nums []int, target int) int {
	l := len(nums)

	if l <= 0 {
		return 0
	}

	if target <= nums[0] {
		return 0
	}

	if target > nums[l-1] {
		return l
	}

	st, ed := 0, l-1
	for st+1 < ed {
		if nums[st] == target {
			return st
		}

		if nums[ed] == target {
			return ed
		}

		mid := (st + ed) / 2
		if nums[mid] == target {
			return mid
		}
		if mid == st || mid == ed {
			break
		}

		if nums[mid] > target {
			ed = mid
			continue
		}

		st = mid
	}

	if nums[st] == target {
		return st
	}
	if nums[ed] == target {
		return ed
	}

	ed = max(st, ed)
	ed = min(ed, l-1)
	return ed
}
