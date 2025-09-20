package roundzero

func rangeBitwiseAnd(left int, right int) int {

	res := helper201(left)
	for i := left + 1; i <= right; i++ {
		if !helper201V1(&res, i) {
			return 0
		}
	}
	return cla(res)
}

func helper201V1(nums *[]bool, n int) bool {
	cnt := 0
	for i := range 32 {

		if n%2 == 1 && (*nums)[i] {
			(*nums)[i] = true
			cnt++
			continue
		} else {
			(*nums)[i] = false
		}

		n = n >> 1
	}

	return cnt != 0
}

func helper201(n int) []bool {
	nums := make([]bool, 32)
	for i := range 32 {

		if n%2 == 1 {
			nums[i] = true
		}

		n = n >> 1
	}
	return nums
}

func cla(nums []bool) int {
	w := 1
	res := 0
	for i := 0; i < 32; i++ {
		if nums[i] {
			res += w
		}
		w = w << 1
	}
	return res
}
