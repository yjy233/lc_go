package roundzero

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {

	st := make([]string, 0, len(tokens))
	st, tokens = tokens, st

	push := func(c string) {
		tokens = append(tokens, c)
	}

	pop := func() (string, error) {
		if len(tokens) <= 0 {
			return "", fmt.Errorf("empty")
		}

		v := tokens[len(tokens)-1]
		tokens = tokens[:len(tokens)-1]
		return v, nil
	}

	calc := func(c string) error {
		secondStr, err := pop()
		if err != nil {
			return err
		}
		firstStr, err := pop()
		if err != nil {
			return err
		}
		second, err := strconv.ParseInt(secondStr, 10, 64)
		if err != nil {
			return err
		}
		first, err := strconv.ParseInt(firstStr, 10, 64)
		if err != nil {
			return err
		}

		res := int64(0)
		switch c {
		case "+":
			res = first + second
		case "-":
			res = first - second
		case "/":
			res = first / second
		case "*":
			res = first * second
		}

		push(strconv.FormatInt(res, 10))
		return nil
	}

	for _, c := range st {

		switch c {
		case "+", "-", "/", "*":
			if err := calc(c); err != nil {
				return -1
			}

		default:
			push(c)
		}
	}

	resStr, err := pop()
	if err != nil {
		return -1
	}

	res, err := strconv.ParseInt(resStr, 10, 64)
	if err != nil {
		return -1
	}
	return int(res)
}
