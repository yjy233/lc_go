package roundzero

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	terms := strings.Split(s, " ")
	words := make([]string, 0, len(terms))

	for _, term := range terms {
		if term == "" || term == " " {
			continue
		}

		words = append(words, term)
	}

	slices.Reverse(words)
	return strings.Join(words, " ")
}
