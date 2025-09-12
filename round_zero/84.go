package roundzero

// lc84
func largestRectangleArea(heights []int) int {

	l := len(heights)
	st := make([]int, 0, l)

	push := func(x int) {
		st = append(st, x)
	}

	pop := func() int {
		v := st[len(st)-1]
		st = st[0 : len(st)-1]
		return v
	}

	getLast := func() int {
		return st[len(st)-1]
	}

	left := make([]int, l)
	right := make([]int, l)

	// left
	for i := 0; i < l; i++ {
		if len(st) == 0 || heights[i] >= heights[getLast()] {
			push(i)
			continue
		}

		for len(st) > 0 && heights[i] < heights[getLast()] {
			v := pop()
			left[v] = i - v
		}

		push(i)
	}

	for len(st) > 0 {
		v := pop()
		left[v] = l - v
	}

	for i := l - 1; i >= 0; i-- {
		if len(st) == 0 || heights[i] >= heights[getLast()] {
			push(i)
			continue
		}

		for len(st) > 0 && heights[i] < heights[getLast()] {
			v := pop()
			right[v] = v - i
		}

		push(i)
	}

	for len(st) > 0 {
		v := pop()
		right[v] = v + 1
	}

	res := 0
	for i := 0; i < l; i++ {
		res = max((left[i]+right[i]-1)*heights[i], res)
	}
	return res
}
