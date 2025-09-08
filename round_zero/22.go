package roundzero

type Item struct {
	Str   []byte
	Left  int
	Right int
}

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}

	res := []Item{
		{
			Str:   []byte{},
			Left:  0,
			Right: 0,
		},
	}

	for i := 0; i < 2*n; i++ {
		tmp := make([]Item, 0)

		for _, item := range res {
			if item.Right > item.Left {
				continue
			}

			if item.Right < item.Left {
				nStr := make([]byte, len(item.Str))
				copy(nStr, item.Str)
				nStr = append(nStr, ')')
				newItem := Item{
					Str:   nStr,
					Left:  item.Left,
					Right: item.Right + 1,
				}

				tmp = append(tmp, newItem)
			}

			if item.Left < n {
				item.Left += 1
				item.Str = append(item.Str, '(')
				tmp = append(tmp, item)
			}
		}

		res = tmp
	}

	realRes := make([]string, 0, len(res))
	for _, item := range res {
		realRes = append(realRes, string(item.Str))
	}

	return realRes
}
