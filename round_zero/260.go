package roundzero

func singleNumber260(nums []int) []int {

	x := 0
	for _, i := range nums {
		x ^= i
	}

	filter := (x) & (-x)

	a, b := 0, 0
	for _, i := range nums {
		if filter&i > 0 {
			a ^= i
		} else {
			b ^= i
		}
	}
	return []int{a, b}
}
