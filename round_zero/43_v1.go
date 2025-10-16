package roundzero

import (
	"fmt"
	"slices"
)

func multiplyV1(num1 string, num2 string) string {
	//n0 := parseIntStr(num1)
	//n1 := parseIntStr(num2)

	res := "0"

	for i := 0; i < len(num2); i++ {
		tmpRes := "0"

		cnt := int(int(num2[i]) - int('0'))
		for c := 0; c < cnt; c++ {
			tmpRes = bigAddV1(tmpRes, num1)
		}
		fmt.Println(tmpRes)
		break
		//for c := 0; c < len(n1)-i-1; c++ {
		//	tmpRes = append(tmpRes, 0)
		//}

		//res = bigAddV1(res, tmpRes)
	}

	fmt.Println(res)
	return ""
}

// 123
// 234
func bigAddV1(s0, s1 string) string {
	n0 := parseIntStr(s0)
	n1 := parseIntStr(s1)
	slices.Reverse(n0)
	slices.Reverse(n1)

	n2 := make([]int, 0, max(len(n1), len(n0))+1)

	jin := 0
	for i := 0; i < max(len(n0), len(n1)); i++ {
		h := jin
		if i < len(n1) {
			h += n1[i]
		}
		if i < len(n0) {
			h += n0[i]
		}

		n2 = append(n2, h%10)
		jin = h / 10
	}

	if jin > 0 {
		n2 = append(n2, jin)
	}
	slices.Reverse(n2)

	r := make([]byte, len(n2))
	for i := range len(n2) {
		r[i] = byte(n2[i] + int('0'))
	}

	return string(r)
}

func parseIntStr(num string) []int {

	l := len(num)

	nInt := make([]int, l)
	for i := range l {
		nInt[i] = int(num[i]) - int('0')
	}

	return nInt
}
