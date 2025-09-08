package roundzero

func isValid(s string) bool {

	cStack := make([]rune, 0)

	popF := func() (bool, rune) {
		l := len(cStack)
		if l == 0 {
			return false, 'a'
		}
		c := cStack[l-1]
		cStack = cStack[:l-1]
		return true, c
	}

	for _, c := range s {
		switch c {
		case ')', ']', '}':
			ok, lc := popF()
			if !ok ||
				(c == ')' && lc != '(') ||
				(c == ']' && lc != '[') ||
				(c == '}' && lc != '{') {
				return false
			}

		default:
			cStack = append(cStack, c)
		}
	}

	return len(cStack) == 0
}
