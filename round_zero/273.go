package roundzero

var (
	wes = []string{
		"Billion",
		"Million",
		"Thousand",
		"Hundred",
	}
	ws = []int{
		1000000000,
		1000000,
		1000,
		100,
	}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	return numberToWordsR(num)
}

func numberToWordsR(num int) string {

	if num <= 9 {
		return getf(num)
	}

	if num < 100 {
		return getTwoWeiN(num)
	}

	for i := 0; i < len(ws); i++ {
		if num < ws[i] {
			continue
		}

		n := num / ws[i]

		if num%ws[i] == 0 {
			return numberToWordsR(n) + " " + wes[i]
		}

		res := numberToWordsR(n) + " " + wes[i] + " " + numberToWordsR(num%ws[i])
		return res

	}
	return ""
}

func getf(num int) string {
	switch num {
	case 0:
		return ""
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	default:
		return ""
	}
}

var (
	tenM = map[int]string{
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
		20: "Twenty",
	}

	shiM = map[int]string{
		1: "Ten",
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
	}
)

func getTwoWeiN(n int) string {
	if n <= 9 {
		return getf(n)
	}

	if n <= 20 {
		return tenM[n]
	}

	if n%10 == 0 {
		return shiM[n/10]
	}

	return shiM[n/10] + " " + getf(n%10)
}
