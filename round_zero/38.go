package roundzero

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	str := make([]byte, 0, 1)
	str = append(str, '1')
	if n <= 1 {
		return string(str)
	}

	for ; n > 1; n-- {
		fmt.Println(n)
		pre := str[0]
		cnt := 1

		tmpStr := make([]byte, 0, 2*len(str))
		for i := 1; i < len(str); i++ {

			if str[i] == pre {
				cnt++
				continue
			}

			tmpStr = append(tmpStr, []byte(strconv.Itoa(cnt))...)
			tmpStr = append(tmpStr, pre)
			pre = str[i]
			cnt = 1
		}

		tmpStr = append(tmpStr, []byte(strconv.Itoa(cnt))...)
		tmpStr = append(tmpStr, pre)

		str = tmpStr

	}

	return string(str)
}
