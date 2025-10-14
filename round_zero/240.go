package roundzero

func searchMatrixV2(matrix [][]int, target int) bool {
	m := len(matrix)
	if m <= 0 {
		return false
	}
	n := len(matrix[0])
	if n <= 0 {
		return false
	}

	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	l := lowerBoundInd(matrix, target, m)
	//fmt.Println(l)
	for i := l; i >= 0; i-- {
		//fmt.Println(i, matrix[i])
		if target < matrix[i][0] || target > matrix[i][n-1] {
			continue
		}
		if binS(matrix[i], target, n) {
			return true
		}
	}

	return false
}

func binS(nums []int, target, n int) bool {

	if target < nums[0] || target > nums[n-1] {
		return false
	}

	l, r := 0, n-1
	for l <= r {
		if nums[r] == target || nums[l] == target {
			return true
		}

		mid := (r + l) / 2
		if nums[mid] == target {
			return true
		}

		if mid == r || l == mid {
			return nums[mid] == target
		}

		if nums[mid] > target {
			r = mid
			continue
		}

		l = mid
	}

	return nums[r] == target || nums[l] == target
}

func lowerBoundInd(matrix [][]int, target int, m int) int {
	l := 0
	r := m - 1

	if target <= matrix[0][0] {
		return 0
	}

	if target >= matrix[m-1][0] {
		return m - 1
	}

	for l <= r {
		if target <= matrix[l][0] {
			return l
		}

		if target >= matrix[r][0] {
			break
		}

		mid := (r + l) / 2
		if mid == r || mid == l {
			break
		}

		if matrix[mid][0] < target {
			l = mid
			continue
		}

		r = mid
	}

	if l > r {
		l, r = r, l
	}

	if matrix[r][0] <= target {
		return r
	}

	return l
}
