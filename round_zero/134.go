package roundzero

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)

	res := 0
	tmp := 0
	idx := 0
	for i := range gas {

		res += gas[i]
		res -= cost[i]

		tmp += gas[i]
		tmp -= cost[i]

		if tmp < 0 {
			tmp = 0

			idx = (i + 1) % l
		}
	}

	if res < 0 {
		return -1
	}
	return idx

}
