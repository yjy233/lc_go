package roundzero

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	singleRowSize := numRows*2 - 2

	tmp := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		tmp[i] = make([]byte, 0)
	}

	for i := 0; i < len(s); i++ {
		rowInd := i % singleRowSize
		if rowInd >= numRows {
			rowInd = numRows - 1 - (rowInd - numRows + 1)
		}

		tmp[rowInd] = append(tmp[rowInd], s[i])
	}

	res := ""

	for i := 0; i < len(tmp); i++ {
		res += string(tmp[i])
	}
	return res
}
