package roundzero

import "math"

var (
	rm map[int]int = map[int]int{}
)

func numSquares(n int) int {
	if n <= 0 {
		return 0
	}

	if nn, ok := rm[n]; ok {
		return nn
	}

	if n == 1 {
		return 1
	}

	mF := math.Sqrt(float64(n))
	nF := int(mF)

	if nF*nF == n {
		return 1
	}

	res := n
	for i := nF; i >= 1; i-- {
		tmp := numSquares(n - i*i)
		if tmp+1 < res {
			res = tmp + 1
		}
	}
	rm[n] = res
	return res
}
