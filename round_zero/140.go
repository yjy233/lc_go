package roundzero

import (
	"fmt"
	"strings"
)

func wordBreakV2(s string, wordDict []string) []string {
	path := make([]string, 0, len(s))
	res := make([]string, 0, len(s))

	mw := len(wordDict[0])
	for _, w := range wordDict {
		if len(w) > mw {
			mw = len(w)
		}
	}
	var dfs func(pos int)
	dfs = func(pos int) {
		if pos >= len(s) {
			tmp := strings.Join(path, " ")
			//fmt.Println()
			res = append(res, tmp)
			return
		}

		for i := pos + 1; i <= min(len(s), i+mw); i++ {
			for _, w := range wordDict {
				fmt.Println(w, s[pos:i])
				if s[pos:i] == w {
					//fmt.Println("---------")
					path = append(path, w)
					dfs(i)
					path = path[:len(path)-1]
				}
			}
		}
	}

	dfs(0)
	return res
}
