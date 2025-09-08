package roundzero

import "sort"

func merge(intervals [][]int) [][]int {
	l := len(intervals)
	if l <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0, l)
	st := intervals[0][0]
	ed := intervals[0][1]

	for i := 1; i < l; i++ {
		if ed >= intervals[i][0] {
			ed = max(ed, intervals[i][1])
		} else {
			res = append(res, []int{st, ed})
			st = intervals[i][0]
			ed = intervals[i][1]
		}
	}

	res = append(res, []int{st, ed})
	return res
}
