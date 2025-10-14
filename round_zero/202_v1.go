package roundzero

var (
	memory202 map[int]bool = make(map[int]bool)
)

func isHappyV32(n int) bool {
	if n == 1 {
		return true
	}

	realRes := true
	path := make(map[int]bool)
	path[n] = true

	for true {
		//fmt.Println(n)
		if r, ok := memory202[n]; ok {
			return r
		}

		res := 0

		for n > 0 {
			w := n % 10
			n /= 10

			res += w * w

		}

		if res == 1 {
			break
		}

		if path[res] {
			realRes = false
			break
		}

		path[res] = true
		n = res
	}

	for k := range path {
		memory202[k] = realRes
	}

	return realRes
}
