package roundzero

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	if newInterval[1] < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}

	if newInterval[0] > intervals[len(intervals)-1][1] {
		return append(intervals, newInterval)
	}

	lb := lowerBound(intervals, newInterval)
	//fmt.Println(lb)
	res := make([][]int, 0, len(intervals)+1)
	for i := lb; i >= 0; i-- {
		a := newInterval
		b := intervals[i]

		if a[0] > b[0] {
			a, b = b, a
		}
		if b[0] > a[1] {
			res = append(res, intervals[:i+1]...)
			//res = append(res, newInterval)
			break
		}

		newInterval[0] = a[0]
		newInterval[1] = max(a[1], b[1])
	}

	//fmt.Println("kkk ",res)
	merged := false
	for i := lb + 1; i < len(intervals); i++ {
		a := newInterval
		b := intervals[i]

		if a[0] > b[0] {
			a, b = b, a
		}
		if b[0] > a[1] {
			//fmt.Println("innter ", newInterval)
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			merged = true
			break
		}

		newInterval[0] = a[0]
		newInterval[1] = max(a[1], b[1])
	}

	if !merged {
		res = append(res, newInterval)
	}

	return res
}

func lowerBound(intervals [][]int, newInterval []int) int {

	a := newInterval[0]

	if a <= intervals[0][0] {
		return 0
	}

	if a >= intervals[len(intervals)-1][0] {
		return len(intervals) - 1
	}

	l, r := 0, len(intervals)-1

	for l <= r {
		if a >= intervals[r][0] {
			return r
		}

		if a <= intervals[l][0] {
			return l
		}

		mid := (r + l) / 2
		if mid == r || mid == l {
			break
		}

		if intervals[mid][0] > a {
			r = mid
			continue
		}

		l = mid
	}

	if l > r {
		r, l = l, r
	}

	//fmt.Println(l,r)
	if intervals[l][0] >= a {
		if intervals[l][0] == a {
			return l
		}
		return l - 1
	}

	if intervals[r][0] <= a {
		if intervals[l][0] == a {
			return r
		}
	}

	return r - 1
}
