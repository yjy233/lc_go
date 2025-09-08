package roundzero

func trap(height []int) int {
	res := 0

	st := make([]int, 0)
	push := func(i int) {
		st = append(st, i)
	}

	pop := func() int {
		v := st[len(st)-1]
		st = st[:len(st)-1]
		return v
	}

	getLast := func() int {
		return st[len(st)-1]
	}

	isEmpty := func() bool {
		return len(st) <= 0
	}

	for i := 0; i < len(height); i++ {
		for !isEmpty() && height[getLast()] <= height[i] {
			bottom := height[pop()]
			if isEmpty() {
				break
			}

			h := max(0, min(height[i], height[getLast()])-bottom)
			w := i - getLast() - 1

			res += h * w
		}

		push(i)
	}

	return res
}
