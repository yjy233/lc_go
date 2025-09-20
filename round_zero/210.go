package roundzero

func findOrder(numCourses int, prerequisites [][]int) []int {
	res := make([]int, 0, numCourses)

	use := make([]bool, numCourses)
	for i := range use {
		use[i] = true
	}

	for _, pair := range prerequisites {
		src := pair[0]
		use[src] = false
	}

	for i := range use {
		if use[i] {
			res = append(res, i)
		}
	}

	if len(res) == 0 {
		return []int{}
	}

	run := true
	for run {
		run = false

	outterCycle:
		for i := range numCourses {
			if use[i] {
				continue
			}

			for _, pair := range prerequisites {
				src, dst := pair[0], pair[1]
				if src != i {
					continue
				}

				if !use[dst] {
					continue outterCycle
				}
			}

			use[i] = true
			res = append(res, i)
			run = true
		}

	}

	for i := range use {
		if !use[i] {
			return []int{}
		}
	}

	return res
}
