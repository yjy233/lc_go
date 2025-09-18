package roundzero

func minCut(s string) int {
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
	/*
	   for _, i := range dp {
	       for _, j := range i{
	           if j {
	               fmt.Print(1, " ")
	           }else {
	               fmt.Print(0, " ")
	           }
	       }
	       fmt.Println("")
	   }*/
	res := make([]int, l)
	for i := range l {
		res[i] = -1
	}

	var dfs func(int, int)
	dfs = func(pos, pre int) {
		if res[l-1] != -1 && pre >= res[l-1] {
			return
		}
		if pos >= l {
			return
		}

		//if res[pos] != -1 && res[pos] <= pre {
		//	return
		//}
		if res[pos] == -1 {
			res[pos] = pre
		} else {
			res[pos] = min(pre, res[pos])
		}

		for i := l - 1; i >= pos; i-- {
			if dp[pos][i] {
				if res[i] == -1 {
					res[i] = pre
				} else {
					if i != pos && res[i] <= pre {
						continue
					}
					res[i] = min(res[i], pre)
				}

				dfs(i+1, pre+1)
			}
		}
	}
	dfs(0, 0)
	//fmt.Println(res)
	return res[l-1]
}
