package roundzero

func isMatchV4(s string, p string) bool {
	l1 := len(s)
	l2 := len(p)

	dp := make([][]bool, l1+1)
	for i := range dp {
		dp[i] = make([]bool, l2+1)
	}

	dp[0][0] = true

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if dp[i+1][j+1] {
				continue
			}

			if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
				continue
			}

			if p[j] != '*' || j < 1 {
				continue
			}

			c := p[j-1]

			if j == 1 {
				dp[0][1] = true
			}
			// a
			// ab*
			// dp[1][1] dp[1][3] j= 2
			if dp[i][j-1] {
				dp[i][j+1] = true
			}
			if dp[i+1][j-1] {
				dp[i+1][j+1] = true
			}

			if c != '.' {
				if s[i] == c && (dp[i][j-1] || dp[i][j+1]) {
					dp[i+1][j+1] = true
				}
				//for k := i; k < l1; k++ {
				//	if s[k] != c {
				//		break
				//	}

				//	dp[k+1][j+1] = true
				//}
				continue
			}

			if dp[i][j-1] || dp[i][j+1] {
				dp[i+1][j+1] = true
			}
			/*
				for k := 0; k <= l1; k++ {
					if dp[k][j-1] {
						for u := k; u <= l1; u++ {
							dp[u][j+1] = true
						}
					}
				}*/
		}
	}

	/*
	   for i := range dp {
	       for j := range dp[i]{
	           if dp[i][j] {
	               fmt.Print(" 1")
	               continue
	           }
	           fmt.Print(" 0")
	       }
	       fmt.Println("")
	   }*/

	return dp[l1][l2]
}
