package roundzero

import "strconv"

func calculate(s string) int {
	st := make([]bool, 0, len(s))
	push := func(x bool) {
		st = append(st, x)
	}
	pop := func() bool {
		v := st[len(st)-1]
		st = st[:len(st)-1]

		return v
	}

	getLast := func() bool {
		return st[len(st)-1]
	}

	push(true)

	tokens := parse224(s)
	//fmt.Println(tokens)

	i := 0
	res := 0
	for i < len(tokens) {
		//fmt.Println(res, i, tokens[i], st)
		if tokens[i] == "(" {
			if i <= 0 {
				push(true)
				i++
				continue
			}

			if getLast() {
				push(tokens[i-1] != "-")
			} else {
				push(tokens[i-1] == "-")
			}

			i++
			continue

		}

		if tokens[i] == ")" {
			pop()
			i++
			continue
		}

		if tokens[i] == "+" || tokens[i] == "-" {

			f := (tokens[i] == "+")
			lf := getLast()

			n, err := strconv.Atoi(tokens[i+1])
			//fmt.Println(n, err)
			if err != nil {
				i++
				continue
			}
			//fmt.Println(f, lf, n)
			if f == lf {
				res += n
			} else {
				res -= n
			}

			i += 2
			continue
		}

		n, _ := strconv.Atoi(tokens[i])
		if getLast() {
			res += n
		} else {
			res -= n
		}

		i++
	}
	return res
}

func parse224(s string) []string {
	l := len(s)

	tokens := make([]string, 0, l)

	pre := 0

	for i := range s {
		switch s[i] {
		case '(', ')', '+', '-':
			if i > pre {
				tokens = append(tokens, s[pre:i])
			}
			tokens = append(tokens, string(s[i]))
			pre = i + 1
		case ' ':
			if i > pre {
				tokens = append(tokens, s[pre:i])
			}
			pre = i + 1
		default:
			continue
		}
	}

	if pre < l {
		tokens = append(tokens, s[pre:])
	}

	return tokens
}
