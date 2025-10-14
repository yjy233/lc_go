package roundzero

import "fmt"

func robV2(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return max(nums[1], nums[0])
	}

	tmp := make([]int, l-1)
	copy(tmp, nums)

	return max(helper213(tmp), helper213(nums[1:]))

}

func helper213(nums []int) int {
	fmt.Println(nums)
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			nums[i] = max(nums[i-1], nums[i])
			continue
		}
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}
	fmt.Println(nums)
	return max(nums[len(nums)-1], nums[len(nums)-2])
}
