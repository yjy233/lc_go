package roundzero

import "strconv"

/*
很简单，就是多叉树分支
*/
func getPermutation(n int, k int) string {

	numsList := make([]int, n)
	numsList[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		numsList[i] = numsList[i+1] * (n - i)
	}

	use := make([]bool, n)

	//n = 4, k = 9
	// 输出："2314"

	res := ""
	for i := 1; i < n; i++ {
		cnt := 0
		for k > numsList[i] {
			k -= numsList[i]
			cnt++
		}

		//fmt.Println(cnt, "xwefwef")

		ji := -1
		for j := 0; j < n; j++ {
			if !use[j] {
				ji++
			}

			if ji == cnt {
				use[j] = true

				res += strconv.FormatInt(int64(j+1), 10)
				break
			}
		}

		if k <= 0 {
			break
		}
	}

	for i := 0; i < n; i++ {
		if !use[i] {
			res += strconv.FormatInt(int64(i+1), 10)
		}

	}

	return res
}
