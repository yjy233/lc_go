package roundzero

import "strings"

func lengthOfLastWord(s string) int {

	terms := strings.Split(s, " ")

	for i := len(terms) - 1; i >= 0; i-- {
		if terms[i] != "" && terms[i] != " " {
			return len(terms[i])
		}
	}

	return 0
}
