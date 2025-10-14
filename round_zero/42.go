package roundzero

import "fmt"

func trap(height []int) int {
	st := make([]int, 0, len(height))

	push := func(v int) {
		st = append(st, v)
	}

	pop := func() int {
		if len(st) <= 0 {
			return -1
		}

		v := st[len(st)-1]
		st = st[:len(st)-1]
		return v
	}

	empty := func() bool {
		return len(st) == 0
	}
	getLast := func() int {
		if empty() {
			return -1
		}
		return st[len(st)-1]
	}

	res := 0
	i := 0
	for i < len(height) {
		fmt.Println(i, res)
		for !empty() && height[getLast()] < height[i] {
			low := pop()
			if empty() {
				break
			}

			l := getLast()
			h := min(height[l], height[i]) - height[low]
			//fmt.Println(i, i-l, h, "----")
			res += max(0, h*(i-l-1))
		}

		push(i)
		i++
	}
	return res
}
