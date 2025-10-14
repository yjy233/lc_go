package roundzero

func findMin(nums []int) int {
	length := len(nums)

	if length <= 0 {
		return 0
	}

	res := nums[0]

	l := 0
	r := len(nums) - 1

	mid := (l + r) / 2
	if mid == l || mid == r {
		return min(nums[l], nums[r])
	}

	if nums[mid] > nums[l] {
		res = min(res, nums[l])
		if mid+1 < length {
			res = min(res, findMin(nums[mid+1:r+1]))
		}
		return res
	}

	if nums[mid] < nums[r] {
		res = min(res, nums[mid])
		if mid-1 > l {
			res = min(res, findMin(nums[l:mid]))
		}
		return res
	}

	if mid < length {
		res = min(res, findMin(nums[mid:]))
	}
	if mid > l {
		res = min(res, findMin(nums[:mid+1]))
	}

	return res
}
