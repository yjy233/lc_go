package roundzero

func generate(numRows int) [][]int {

	res := make([][]int, 0, numRows)
	if numRows < 1 {
		return res
	}

	res = append(res, []int{1})

	for n := 1; n < numRows; n++ {
		num := n + 1

		row := make([]int, num)
		row[0], row[num-1] = 1, 1

		for i := 1; i < num-1; i++ {
			row[i] = res[n-1][i-1] + res[n-1][i]
		}

		res = append(res, row)
	}

	return res
}
