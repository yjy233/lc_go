package roundzero

func candy(ratings []int) int {

	l := len(ratings)
	left := make([]int, l)
	right := make([]int, l)

	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	for i := l - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	res := l
	for i := 0; i < l; i++ {
		res += max(left[i], right[i])
	}

	return res
}
