package roundzero

import "fmt"

func multiply(num1 string, num2 string) string {

	if len(num2) <= 0 || num1 == "0" || num2 == "0" {
		return "0"
	}

	singN := num2[0]

	baseRes := "0"
	//fmt.Println(int(singN - '0'))
	for i := 0; i < int(singN-'0'); i++ {
		//fmt.Println(baseRes)
		baseRes = bigAdd(baseRes, num1)
	}
	//fmt.Println(",,,,", baseRes)
	num2 = num2[1:]

	if len(num2) <= 0 {
		return baseRes
	}

	for i := 0; i < len(num2); i++ {
		baseRes += "0"
	}
	//fmt.Println("mmmmm", baseRes)
	return bigAdd(baseRes, multiply(num1, num2))

}

func bigAdd(nums1 string, nums2 string) string {
	//if nums1 == "0" || nums2 == "0" {
	//	return "0"
	//}

	intL1 := parseInt(nums1)
	intL2 := parseInt(nums2)
	fmt.Println(intL1, intL2)

	l := max(len(intL1), len(intL2))
	resIntL := make([]int, l, l+1)

	jin := 0
	for i := 0; i < l; i++ {
		he := jin
		if len(intL1) > i {
			he += intL1[i]
		}
		if len(intL2) > i {
			he += intL2[i]
		}

		resIntL[i] = he % 10
		jin = he / 10
	}

	if jin > 0 {
		resIntL = append(resIntL, jin)
	}

	resByte := make([]byte, len(resIntL))
	for ind, i := range resIntL {
		//fmt.Println("----",i)
		resByte[len(resByte)-1-ind] = byte(i + '0')
	}

	//fmt.Println("add res ", string(resByte))
	return string(resByte)

}

func parseInt(nums1 string) []int {
	l := make([]int, len(nums1))

	for i := len(nums1) - 1; i >= 0; i-- {
		l[len(nums1)-1-i] = int(nums1[i] - '0')
	}

	return l
}
