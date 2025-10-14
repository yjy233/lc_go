package roundzero

func isMatchV11(s string, p string) bool {

	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j <= len(p); j++ {
			if dp[i][j] {
				continue
			}

			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			if p[j-1] != '*' || j < 2 {
				continue
			}
			//fmt.Println("fewfewfwef")

			if dp[i][j-2] {
				dp[i][j] = true
			}

			if dp[i-1][j-2] {
				dp[i-1][j] = true
			}

			c := p[j-2]
			//fmt.Println("sss ",string(c))
			if c != '.' {
				if (j == 2 && i == 1) || dp[i-1][j-2] {
					for k := max(1, i); k <= len(s); k++ {
						if s[k-1] != c {
							break
						}
						dp[k][j] = true
					}
				}

				continue
			}

			//fmt.Println("ww ",i,j)
			if j == 2 || dp[i-1][j-2] {
				for k := max(1, i); k <= len(s); k++ {
					dp[k][j] = true
				}
			}

		}
	}

	/*
	   for i := range dp{
	       for _,j := range dp[i]{
	           if j {
	               fmt.Print(" 1")
	           }else {
	               fmt.Print(" 0")
	           }
	       }
	       fmt.Println("")
	   }*/
	return dp[len(s)][len(p)]
}
