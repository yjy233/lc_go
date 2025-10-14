package roundzero

func generateParenthesis(n int) []string {
	type item struct {
		l   int
		r   int
		str []byte
	}

	copyItem := func(i *item) *item {

		byteL := make([]byte, len(i.str), 2*n)
		copy(byteL, i.str)

		return &item{
			l:   i.l,
			r:   i.r,
			str: byteL,
		}
	}

	res := make([]*item, 0, 2*n)
	tmp := make([]*item, 0, 2*n)

	res = append(res, &item{
		str: make([]byte, 0, 2*n),
	})

	for i := 0; i < 2*n; i++ {
		for j := range res {
			oldItem := res[j]

			if oldItem.r < n && oldItem.r < oldItem.l {
				newItem := copyItem(oldItem)
				newItem.r++
				newItem.str = append(newItem.str, ')')

				tmp = append(tmp, newItem)
			}

			if oldItem.l < n {
				oldItem.l++
				oldItem.str = append(oldItem.str, '(')

				tmp = append(tmp, oldItem)
			}
		}

		res, tmp = tmp, res
		tmp = tmp[:0]
	}

	realRes := make([]string, 0, len(res))
	for _, item := range res {
		realRes = append(realRes, string(item.str))
	}
	return realRes

}
