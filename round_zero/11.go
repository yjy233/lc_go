package roundzero

func maxArea(height []int) int {
	res := 0

	left, right := 0, len(height)-1
	for left < right {
		res = max(res, min(height[left], height[right])*(right-left))

		if height[left] < height[right] {
			left++
			continue
		}

		right--
	}

	return res
}
