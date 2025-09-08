package roundzero

import "math"

func myPow(x float64, n int) float64 {
	isPos := n >= 0
	n = int(math.Abs(float64(n)))
	if n == 0 {
		return 1
	}

	absA := helperMyPow(x, n)
	if isPos {
		return absA
	}
	return 1 / absA
}

func helperMyPow(x float64, n int) float64 {

	if n == 1 {
		return x
	}

	if n%2 == 0 {
		return helperMyPow(x*x, n/2)
	}
	return helperMyPow(x*x, n/2) * x
}
