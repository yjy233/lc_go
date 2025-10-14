package roundzero

type preTree struct {
	Stop bool
	Str  string
	Next [26]*preTree
}

func findWords(board [][]byte, words []string) []string {

	root := &preTree{}

	m := len(board)
	if m == 0 {
		return nil
	}

	n := len(board[0])
	if n == 0 {
		return nil
	}

	path := make([][]bool, m)
	for i := range path {
		path[i] = make([]bool, n)
	}

	buildHelper := func(word string) {

		tmpR := root
		for _, i := range word {
			offset := int(i) - int('a')

			if tmpR.Next[offset] == nil {
				tmpR.Next[offset] = &preTree{}
			}
			tmpR = tmpR.Next[offset]
		}

		tmpR.Str = word
		//fmt.Println(word)
		tmpR.Stop = true
	}

	for _, w := range words {
		buildHelper(w)
	}

	res := make(map[string]bool)

	var dfs func(int, int, *preTree)
	dfs = func(x, y int, r *preTree) {
		//fmt.Println(x,y,string(board[x][y]))
		if len(res) >= len(words) {
			return
		}

		if r == nil {
			return
		}
		c := board[x][y]

		if r.Str != "" && r.Stop {
			res[r.Str] = true
		}
		offset := int(c) - int('a')
		if r.Next[offset] == nil {
			return
		}

		if r.Next[offset].Str != "" && r.Next[offset].Stop {
			res[r.Next[offset].Str] = true
		}

		path[x][y] = true

		if x < m-1 && !path[x+1][y] {
			dfs(x+1, y, r.Next[offset])
		}

		if x > 0 && !path[x-1][y] {
			dfs(x-1, y, r.Next[offset])
		}

		if y < n-1 && !path[x][y+1] {
			dfs(x, y+1, r.Next[offset])
		}

		if y > 0 && !path[x][y-1] {
			dfs(x, y-1, r.Next[offset])
		}

		path[x][y] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, root)
		}
	}

	realRes := make([]string, 0, len(res))
	for k := range res {
		if k == "" {
			continue
		}
		realRes = append(realRes, k)
	}

	//sort.Slice(realRes, func(i,j int)bool {
	//    return realRes[i] <realRes[j]
	//})
	//fmt.Println(res)

	return realRes
}
