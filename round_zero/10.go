package roundzero

import "fmt"

func isMatch(s string, p string) bool {
	sl := len(s)
	pl := len(p)
	if sl <= 0 {
		return true
	}
	if pl <= 0 {
		return false
	}

	dp := make([][]bool, sl+1)
	for i := 0; i <= sl; i++ {
		dp[i] = make([]bool, pl+1)
	}
	dp[0][0] = true

	for i := 0; i < sl; i++ {
		for j := 0; j < pl; j++ {
			if s[i] == p[j] {
				if dp[i][j] {
					dp[i+1][j+1] = true
					continue
				}
			}

			if p[j] == '.' {
				if dp[i][j] {
					dp[i+1][j+1] = true
					continue
				}
			}

			if j <= 0 || p[j] != '*' {
				continue
			}
			//fmt.Println(dp[i][j-1])
			//fmt.Println(dp[i+1][j-1])
			//fmt.Println(i,j)
			//if dp[i+1][j-1] {
			//	dp[i+1][j+1] = true
			//}
			for k := 0; k <= sl; k++ {
				if dp[k][j-1] {
					dp[k][j+1] = true
				}
			}

			c := p[j-1]
			if c == '.' {
				if j == 1 {
					for k := 0; k <= sl; k++ {
						dp[k][j+1] = true
					}

					continue
				}

				if dp[i][j-1] {
					for k := i + 1; k <= sl; k++ {
						dp[k][j+1] = true
					}

					continue
				}
			}

			if j == 1 {
				dp[0][j+1] = true
				for k := 0; k < sl; k++ {
					if s[k] != c {
						break
					}
					dp[k+1][j+1] = true
				}

				continue
			}
			fmt.Println("xwefwe. " + string(c))
			if dp[i][j-1] {
				for k := i; k < sl; k++ {
					if s[k] != c {
						break
					}
					dp[k+1][j+1] = true
				}

			}

		}
	}

	return dp[sl][pl]
}
