package roundzero

func getRow(rowIndex int) []int {

	res := make([]int, rowIndex+1)

	res[0] = 1

	for i := 1; i <= rowIndex; i++ {
		res[i] = res[i-1] * (rowIndex + 1 - i) / i
	}

	return res
}
