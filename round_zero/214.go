package roundzero

import "slices"

func shortestPalindrome(s string) string {
	if s == "" {
		return s
	}
	c := s[0]

	cs := make([]int, 0)
	for i := range s {
		if s[i] == c {
			cs = append(cs, i)
		}
	}

	r := 0
	for i := len(cs) - 1; i >= 0; i-- {
		if judge3(s, cs[i]) {
			r = cs[i]
			break
		}
	}

	sub := ([]byte)(s[r+1:])
	slices.Reverse(sub)

	return string(sub) + s
}

func judge3(s string, r int) bool {
	if r == 0 {
		return true
	}

	r -= 1
	l := 1
	for l <= r {
		if s[l] == s[r] {
			l++
			r--
			continue
		}
		return false
	}
	return true
}
