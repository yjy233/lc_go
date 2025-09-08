package roundzero

func rotate(matrix [][]int) {
	// 1 2 3 4
	// 5 6 7 8
	// 9 1 2 3
	// 4 5 6 7

	l := len(matrix)
	for bias := 0; bias < l/2; bias++ {
		for i := bias; i < l-bias-1; i++ {
			// line bias, l-1-bias
			// col bias l-1-bias

			tmp := matrix[i][l-1-bias]
			//fmt.Println(tmp)
			matrix[i][l-1-bias] = matrix[bias][i]
			//fmt.Println(matrix[i][l-1-bias])
			tmp1 := matrix[l-1-bias][l-1-i]
			//fmt.Println(tmp1, tmp)
			matrix[l-1-bias][l-1-i] = tmp
			tmp = matrix[l-1-i][bias]
			matrix[l-1-i][bias] = tmp1
			matrix[bias][i] = tmp

			//break
		}
	}

}
