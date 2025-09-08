package roundzero

import "strings"

func simplifyPath(path string) string {
	terms := strings.Split(path, "/")

	st := make([]string, 0, len(terms))
	push := func(term string) {
		st = append(st, term)
	}

	pop := func() {
		if len(st) <= 0 {
			return
		}
		st = st[:len(st)-1]
	}

	for _, term := range terms {
		if term == "/" || term == "." || term == " " || term == "" {
			continue
		}

		if term == ".." {
			pop()
			continue
		}

		if term == " " {
			continue
		}
		//fmt.Println(term,"---")
		push(term)
	}

	res := "/"
	if len(st) == 0 {
		return res
	}

	return res + strings.Join(st, "/")
}
