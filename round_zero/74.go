package roundzero

func searchMatrix(matrix [][]int, target int) bool {
	lowerBound := helperLowerBoundSearchMarix(matrix, target)
	if lowerBound == -1 {
		return false
	}
	return binarySearch(matrix[lowerBound], target)
}

func helperLowerBoundSearchMarix(matrix [][]int, target int) int {
	l := len(matrix)

	if matrix[0][0] > target {
		return -1
	}

	if matrix[l-1][0] <= target {
		return l - 1
	}

	l, r := 0, l-1
	for l < r {
		if matrix[l][0] == target {
			return l
		}

		if matrix[r][0] <= target {
			return r
		}

		mid := (l + r) / 2
		if matrix[mid][0] == target {
			return mid
		}

		if mid == l || mid == r {
			break
		}

		if matrix[mid][0] >= target {
			r = mid
			continue
		}
		l = mid + 1
	}

	r = max(r, l)

	for i := r; i >= 0; i-- {
		if matrix[i][0] <= target {
			return i
		}
	}

	return 0
}

func binarySearch(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		if nums[l] == target || nums[r] == target {
			return true
		}

		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}

		if mid == r || mid == l {
			if nums[l] == target || nums[r] == target {
				return true
			}
			return false
		}

		if nums[mid] < target {
			l = mid + 1
			continue
		}

		r = mid

	}

	return false
}
