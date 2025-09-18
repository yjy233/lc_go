package roundzero

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)

	total := 0
	cur := 0
	res := 0
	for i := range l {
		c := gas[i] - cost[i]

		total += c
		cur += c
		if cur < 0 {
			cur = 0
			res = i + 1
		}
	}

	if total < 0 {
		return -1
	}

	return res % l

}
