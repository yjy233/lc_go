package roundzero

import "strconv"

func calculate227(s string) int {
	tokens := parse227(s)
	//fmt.Println(tokens)
	st := make([]int, 0, len(tokens))
	push := func(x int) {
		st = append(st, x)
	}
	pop := func() int {
		v := st[len(st)-1]
		st = st[:len(st)-1]
		return v
	}

	i := 0
	for i < len(tokens) {
		//fmt.Println(st, i)
		if tokens[i] == "+" {
			n, _ := strconv.Atoi(tokens[i+1])
			push(n)
			i += 2
			continue
		}

		if tokens[i] == "-" {
			n, _ := strconv.Atoi(tokens[i+1])
			//fmt.Println("xwefwe ",n, tokens[i+1])
			push(-n)
			i += 2
			continue
		}

		if tokens[i] == "*" || tokens[i] == "/" {
			n, _ := strconv.Atoi(tokens[i+1])
			p := pop()
			c := p * n
			if tokens[i] == "/" {
				c = p / n
			}
			push(c)
			i += 2
			continue
		}

		n, _ := strconv.Atoi(tokens[i])
		i++
		push(n)
	}

	res := 0
	for _, i := range st {
		res += i
	}

	return res
}

func parse227(s string) []string {
	l := len(s)
	tokens := make([]string, 0, l)

	pre := 0

	for i := range s {
		switch s[i] {
		case '+', '-', '/', '*':
			if pre < i {
				tokens = append(tokens, s[pre:i])
			}
			tokens = append(tokens, string(s[i]))
			pre = i + 1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			if pre < i {
				term := s[pre:i]
				if term != " " && term != "" {
					tokens = append(tokens, s[pre:i])
				}
			}
			pre = i + 1
			continue
		}
	}

	if pre < l {
		term := s[pre:]
		if term != " " && term != "" {
			tokens = append(tokens, term)
		}
	}
	return tokens
}
