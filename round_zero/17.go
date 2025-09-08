package roundzero

func letterCombinations(digits string) []string {
	n2charList := map[rune][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	if digits == "" {
		return []string{}
	}

	res := make([][]byte, 0)
	res = append(res, []byte{})
	for _, x := range digits {
		tmpRes := make([][]byte, 0)

		for _, c := range n2charList[x] {
			//fmt.Println(c)
			for _, p := range res {
				tp := make([]byte, len(p))
				copy(tp, p)
				tmpRes = append(tmpRes, append(tp, c))
			}
		}

		res = tmpRes
	}

	sList := make([]string, 0)
	for _, t := range res {
		sList = append(sList, string(t))
	}

	return sList
}
