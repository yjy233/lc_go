package roundzero

func intToRoman(num int) string {
	c2Num := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n2C := map[int]byte{
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}

	cList := []byte{'M', 'C', 'X', 'I'}

	cInd := 0
	resB := make([]byte, 0)
	for num > 0 {
		c := cList[cInd]
		n := c2Num[c]
		if num < n {
			cInd++
			continue
		}

		if c == 'M' {
			num -= n
			resB = append(resB, c)
			continue
		}

		if num/n == 9 {
			resB = append(resB, c)
			resB = append(resB, n2C[n*10])
			num -= 9 * n
			continue
		}

		if num/n >= 5 {
			resB = append(resB, n2C[n*5])
			num -= 5 * n
			continue
		}

		if num/n == 4 {
			resB = append(resB, c)
			resB = append(resB, n2C[n*5])
			num -= 4 * n
			continue
		}
		resB = append(resB, c)
		num -= n
	}

	return string(resB)
}
