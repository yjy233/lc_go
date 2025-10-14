package roundzero

import (
	"slices"
)

func shortestPalindrome(s string) string {
	m := make(map[string]bool)

	for i := 1; i <= len(s); i++ {
		m[s[:i]] = true
	}

	sb := []byte(s)

	slices.Reverse(sb)
	rs := string(sb)

	for i := 0; i < len(s); i++ {
		if m[s[i:]] {
			return string(rs[:i]) + s
		}
	}

	return rs[:len(s)-1] + s
}
