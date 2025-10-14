package roundzero

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	tokens := make([]string, 0, len(nums))
	for _, i := range nums {
		tokens = append(tokens, strconv.Itoa(i))
	}

	sort.Slice(tokens, func(i, j int) bool {
		a := tokens[i] + tokens[j]
		b := tokens[j] + tokens[i]
		l := min(len(a), len(b))

		for k := 0; k < l; k++ {
			if a[k] > b[k] {
				return true
			}

			if a[k] < b[k] {
				return false
			}
		}
		return true
	})

	realTokens := make([]string, 0, len(tokens))
	isFrist := true
	for _, t := range tokens {
		if t == "0" && isFrist {
			continue
		}

		isFrist = false
		realTokens = append(realTokens, t)
	}

	//fmt.Println(tokens)
	res := strings.Join(realTokens, "")
	if res == "" {
		return "0"
	}
	return res
}
