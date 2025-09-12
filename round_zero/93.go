package roundzero

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	stack := make([]string, 0, 4)
	res := make([]string, 0)

	helperRestorIpAddress(s, &stack, 0, &res)
	return res
}

func helperRestorIpAddress(s string, stack *[]string, ind int, res *[]string) {
	if ind >= len(s) {
		return
	}

	for i := ind; i < min(i+3, len(s)); i++ {
		term := s[ind : i+1]
		if !validIPAddress(term) {
			continue
		}

		if i == len(s)-1 {
			if len(*stack) == 3 {
				(*stack) = append((*stack), term)
				(*res) = append((*res), strings.Join((*stack), "."))
				(*stack) = (*stack)[:3]
			}
			continue
		}

		(*stack) = append((*stack), term)
		helperRestorIpAddress(s, stack, i+1, res)
		(*stack) = (*stack)[:len(*stack)-1]
	}
}

func validIPAddress(term string) bool {
	if term[0] == '0' {
		return len(term) == 1
	}

	num, err := strconv.Atoi(term)
	if err != nil {
		return false
	}

	return num >= 0 && num <= 255
}
