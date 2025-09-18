package roundzero

func partition131(s string) [][]string {
	l := len(s)

	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, l)
	}

	for i := range l {
		dp[i][i] = true
	}

	for i := 1; i < l; i++ {
		if s[i] == s[i-1] {
			dp[i-1][i] = true
		}
	}

	for w := 3; w <= l; w++ {
		for i := 0; i+w-1 < l; i++ {
			if s[i] == s[i+w-1] {
				dp[i][i+w-1] = dp[i+1][i+w-2]
			}
		}
	}

	path := make([]string, 0, l)

	res := make([][]string, 0)

	var dfs func(int)
	dfs = func(pos int) {

		if pos >= l {
			return
		}

		for i := pos; i < l; i++ {
			if dp[pos][i] {
				path = append(path, s[pos:i+1])

				if i == l-1 {
					tmp := make([]string, len(path))
					copy(tmp, path)
					res = append(res, tmp)
				}

				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}

	}
	dfs(0)
	return res
}
