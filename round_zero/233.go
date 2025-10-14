package roundzero

import (
	"fmt"
	"math"
	"strconv"
)

func countDigitOne(n int) int {

	s := ([]byte)(fmt.Sprint(n))
	res := 0

	for i := len(s) - 2; i >= 1; i-- {
		fmt.Println(string(s[i]), res)
		right := 1
		if i < len(s)-1 {
			right = 1
			for j := len(s) - 1; j > i; j-- {
				right *= 10
			}
		}

		left := 1
		if i > 0 {
			left, _ = strconv.Atoi(string(s[:i]))
		}
		// 220
		if s[i] >= '2' && i > 0 {
			left++
		}

		fmt.Println("-", left, right)
		res += left * right

		if s[i] == '1' {
			if i == len(s)-1 {
				res++
			} else {
				tmpS := string(s[i+1:])
				tmpC, _ := strconv.Atoi(tmpS)
				res += tmpC
				res++
			}
		}

	}

	fmt.Println(res, "--fwefwef")
	//
	{
		tmpS, _ := strconv.Atoi(string(s[:len(s)-1]))

		res += tmpS
		if s[len(s)-1] >= '1' {
			res++
		}
	}

	if len(s) > 1 {
		if s[0] == '1' {
			tmpC, _ := strconv.Atoi(string(s[1:]))
			res += tmpC
			res++
		} else {
			res += int(math.Pow(10.0, float64(len(s))-1))
		}
	}

	return res
}
