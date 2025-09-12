package roundzero

/*
没什么好说的，死记硬背就是
*/
func grayCode(n int) []int {

	powN := 1 << n
	res := make([]int, powN)
	for i := 0; i < powN; i++ {
		res[i] = i ^ (i >> 1)
	}
	return res
}
