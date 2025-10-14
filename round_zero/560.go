package roundzero

func subarraySum(nums []int, k int) int {

	res := 0
	memory := make(map[int]int)

	tmp := 0
	for i := range nums {
		tmp += nums[i]

		if tmp == k {
			res++
		}

		res += memory[tmp-k]
		memory[tmp] += 1
	}
	return res
}
