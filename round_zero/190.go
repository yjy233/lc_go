package roundzero

import "slices"

func reverseBits(n int) int {
	dp := make([]int, 32)

	for i := 0; i < 32; i++ {
		dp[i] = n % 2
		n = n >> 1
	}
	//fmt.Println(dp)
	slices.Reverse(dp)

	w := 1
	res := 0
	for i := 0; i < 32; i++ {
		if dp[i] == 1 {
			res += w
		}
		w = w << 1
	}

	return res
}
