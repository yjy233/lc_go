package roundzero

func longestPalindromeSubseq(s string) int {
	length := len(s)

	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}

	//printDPF := func() {
	//    for i := range dp{
	//        for j := range dp[i] {
	//            fmt.Print(dp[i][j], " ")
	//        }
	//        fmt.Println("")
	//    }
	//    fmt.Println("-----------------")
	//}

	for i := range dp {
		dp[i][i] = 1
	}

	//printDPF()
	for i := 1; i < length; i++ {
		if s[i] == s[i-1] {
			dp[i-1][i] = 2
			continue
		}
		dp[i-1][i] = 1
	}
	//printDPF()
	// a b a
	for l := 3; l <= length; l++ {
		for i := 0; i+l-1 < length; i++ {
			dp[i][i+l-1] = max(dp[i+1][i+l-2], dp[i][i+l-2], dp[i+1][i+l-1])
			if s[i] == s[i+l-1] {
				dp[i][i+l-1] = max(dp[i][i+l-1], dp[i+1][i+l-2]+2)
				continue
			}

		}

	}
	//printDPF()
	return dp[0][length-1]
}
