package roundzero

func longestConsecutive(nums []int) int {

	res := 1
	if len(nums) <= 0 {
		return 0
	}
	n2Exist := make(map[int]bool)
	for _, n := range nums {
		n2Exist[n] = true
	}

	dedup := make(map[int]int)
	for n := range n2Exist {
		if dedup[n-1] > 0 || dedup[n] > 0 {
			continue
		}

		//fmt.Println(n)
		tmp := n + 1
		for n2Exist[tmp] && dedup[tmp] == 0 {
			tmp++
		}
		res = max(res, tmp-n+dedup[tmp])
		dedup[n] = tmp - n + dedup[tmp]
	}
	return res
}
