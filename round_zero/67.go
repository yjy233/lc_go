package roundzero

func addBinary(a string, b string) string {
	aL := parsBianary(a)
	bL := parsBianary(b)

	res := make([]int, 0, max(len(aL), len(bL))+1)

	jin := 0
	for i := 0; i < max(len(aL), len(bL)); i++ {

		tmp := jin
		jin = 0

		if i < len(aL) {
			tmp += aL[i]
		}
		if i < len(bL) {
			tmp += bL[i]
		}

		res = append(res, tmp%2)
		jin = tmp / 2
	}

	if jin > 0 {
		res = append(res, jin)
	}

	//   fmt.Println(res)
	byteL := make([]byte, len(res))
	for i := len(res) - 1; i >= 0; i-- {

		c := byte('1')
		if res[i] == 0 {
			c = '0'
		}

		byteL[len(res)-1-i] = c
	}

	return string(byteL)
}

func parsBianary(s string) []int {
	l := len(s)
	res := make([]int, 0, l)
	for i := l - 1; i >= 0; i-- {
		if s[i] == '1' {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
	}

	return res
}
