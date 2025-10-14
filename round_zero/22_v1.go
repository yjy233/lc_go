package roundzero

func generateParenthesisV1(n int) []string {

	type Item struct {
		l      int
		r      int
		buffer []byte
	}

	res := make([]*Item, 0, 1<<n)
	tmpRes := make([]*Item, 0, 1<<n)

	res = append(res, &Item{
		buffer: make([]byte, 0, 2*n),
	})

	for c := 0; c < 2*n; c++ {

		for _, item := range res {
			if item.r < item.l {
				newItem := &Item{
					l:      item.l,
					r:      item.r + 1,
					buffer: make([]byte, len(item.buffer), 2*n),
				}
				copy(newItem.buffer, item.buffer)
				newItem.buffer = append(newItem.buffer, ')')
				tmpRes = append(tmpRes, newItem)
			}

			if item.l < n {
				item.l++
				item.buffer = append(item.buffer, '(')
				tmpRes = append(tmpRes, item)
			}
		}

		res, tmpRes = tmpRes, res
		tmpRes = tmpRes[:0]
	}

	realRes := make([]string, 0, len(res))
	for _, item := range res {
		realRes = append(realRes, string(item.buffer))
	}
	return realRes
}
