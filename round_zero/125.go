package roundzero

import (
	"fmt"
	"strings"
)

func isPalindromeV1(s string) bool {

	s = strings.ToLower(s)
	fmt.Println(s)
	validC := func(r byte) bool {
		fmt.Println(r, 'a', 'z')
		return (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
	}

	l := 0
	for i := 0; i < len(s); i++ {
		if validC(s[i]) {
			l = i
			break
		}
	}
	r := l - 1
	for i := len(s) - 1; i >= 0; i-- {
		if validC(s[i]) {
			r = i
			break
		}
	}
	fmt.Println(l, r)
	for l < r {
		fmt.Println(l, r)
		if s[l] != s[r] {
			return false
		}

		l++
		for !validC(s[l]) && l < len(s) {
			l++
		}

		r--
		for !validC(s[r]) && r >= 0 {
			r--
		}
	}

	fmt.Println(l, r)
	return true
}
