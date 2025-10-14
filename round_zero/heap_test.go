package roundzero

import (
	"fmt"
	"testing"
)

func TestHead(t *testing.T) {

	nums := []int{32, 435, 656, 12, 3}

	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Println(i)
		adjustHeap(nums, i)
	}

	fmt.Println(nums)
	newL := make([]int, 0)

	for len(nums) > 0 {
		newL = append(newL, nums[0])

		nums[0] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		adjustHeap(nums, 0)
	}
	newL = append(newL, nums...)
	t.Log(newL)
}
