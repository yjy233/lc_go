package roundzero

/*
二分找到第一个比他小的就行
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	l := len(intervals)
	if l == 0 {
		return [][]int{newInterval}
	}

	//fmt.Println("0----")
	lowerInd := findBiggerBoundInsert(intervals, newInterval[0])
	//fmt.Println("1----")
	res := make([][]int, 0, l+1)

	if lowerInd > 0 {
		res = append(res, intervals[:lowerInd]...)
	}
	//fmt.Println(lowerInd)
	if lowerInd >= 0 {
		r0, _, flap := mergeInsert(newInterval, intervals[lowerInd])
		if !flap {
			res = append(res, r0)
		} else {
			newInterval = r0
		}
	}

	//fmt.Println(res)
	//fmt.Println(newInterval)
	//fmt.Println("--wefewfwe")

	ok := false
	for i := lowerInd + 1; i < l; i++ {
		newInterval, _, flap := mergeInsert(newInterval, intervals[i])

		if !flap {
			ok = true
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			break
		}
	}

	if !ok {
		res = append(res, newInterval)
	}

	return res
}

func findBiggerBoundInsert(intervals [][]int, st int) int {
	l := len(intervals)
	if intervals[0][0] >= st {
		if intervals[0][0] > st {
			return -1
		}
		return 0
	}

	if st >= intervals[l-1][0] {
		return l - 1
	}

	left, right := 0, l-1
	for left < right {
		if intervals[left][0] == st {
			return left
		}

		if intervals[right][0] == st {
			return right
		}

		mid := (left + right) / 2
		if intervals[mid][0] == st {
			return mid
		}

		if mid == left || mid == right {
			if intervals[left][0] == st {
				return left
			}

			if intervals[right][0] == st {
				return right
			}

			if intervals[left][0] <= st && intervals[right][1] >= st {
				return left
			}
		}

		if intervals[mid][0] > st {
			right = mid
			continue
		}
		left = mid + 1
	}

	for i := right; i >= 0; i-- {
		if intervals[i][0] <= st {
			return i
		}
	}

	return 0
}

func mergeInsert(r0, r1 []int) ([]int, []int, bool) {
	if r0[0] >= r1[0] {
		r0, r1 = r1, r0
	}

	if r0[1] < r1[0] {
		return r0, r1, false
	}

	r0[0] = min(r0[0], r1[0])
	r0[1] = max(r0[1], r1[1])

	return r0, nil, true
}
