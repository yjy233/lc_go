package roundzero

import "slices"

func hIndexV2(citations []int) int {
	length := len(citations)

	slices.Reverse(citations)
	l, r := 0, length-1

	if citations[r] >= length {
		return length
	}

	if citations[l] <= 0 {
		return 0
	}

	for l <= r {
		if citations[r] >= length {
			return r
		}

		if citations[l] < l+1 {
			return l
		}

		mid := (l + r) / 2
		if mid == r || mid == l {
			break
		}

		if citations[mid] >= mid+1 {
			l = mid
			continue
		}

		r = mid
	}

	if l > r {
		r, l = l, r
	}
	//fmt.Println(citations)
	//fmt.Println(l,r)

	if citations[r] >= r+1 {
		return r + 1
	}

	return l + 1
}
