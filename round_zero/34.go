package roundzero

import "fmt"

func searchRange(nums []int, target int) []int {

	ind := findFisrtInd(nums, target)
	left := ind
	right := ind

	fmt.Println(ind)
	if ind == -1 {
		return []int{left, right}
	}

	left = findBounder(nums, ind, 0)
	fmt.Println("-----------------")
	right = findBounder(nums, ind, 1)
	return []int{left, right}
}

func findFisrtInd(nums []int, target int) int {
	left, right := 0, len(nums)-1

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

		if nums[mid] > target {
			right = mid
			continue
		}
		left = mid
	}

	return -1
}

/*
@param leftOrRight 左边界 0 右边界 1
*/
func findBounder(nums []int, ind, leftOrRight int) int {

	isLeft := leftOrRight == 0
	left := 0
	right := ind
	if !isLeft {
		left = ind
		right = len(nums) - 1
	}

	for left < right {
		fmt.Println(left, right)
		mid := (left + right) / 2
		if mid == left || right == mid {
			break
		}

		if isLeft {
			if nums[left] == nums[ind] {
				break
			}

			if nums[mid] < nums[ind] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			if nums[right] == nums[ind] {
				break
			}

			if nums[mid] > nums[ind] {
				right = mid - 1
			} else {
				left = mid
			}
		}
	}

	if isLeft {
		if left < right {
			if nums[left] == nums[ind] {
				return left
			}
			return right
		}
		if nums[right] == nums[ind] {
			return right
		}
		return left
	}

	if left < right {
		if nums[right] == nums[ind] {
			return right
		}
		return left
	}
	if nums[left] == nums[ind] {
		return left
	}
	return right
}
