package roundzero

import (
	"container/heap"
)

type hip []int

func (this *hip) Len() int {
	return len(*this)
}

func (this *hip) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

func (this *hip) Less(i, j int) bool {
	return (*this)[i] > (*this)[j]
}

func (this *hip) Push(i any) {

	(*this) = append((*this), i.(int))
}

func (this *hip) Pop() any {

	v := (*this)[len(*this)-1]
	(*this) = (*this)[:len(*this)-1]
	return v
}

func findKthLargest(nums []int, k int) int {
	hp := (hip)(nums)

	go func() {
		for i := 0; i < k; i++ {
			heap.Push(&hp, nums[i])
		}
	}()
	heap.Init(&hp)
	val := 0
	for k > 0 {
		val = heap.Pop(&hp).(int)
		k--
	}

	//runtime.GOMAXPROCS()
	return val
}
