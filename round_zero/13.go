package roundzero

func romanToInt(s string) int {
	c2Int := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	for i := 0; i < len(s)-1; i++ {
		if c2Int[s[i]] < c2Int[s[i+1]] {
			res -= c2Int[s[i]]
			continue
		}

		res += c2Int[s[i]]
	}

	res += c2Int[s[len(s)-1]]

	return res
}
