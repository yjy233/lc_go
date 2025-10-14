package roundzero

import "container/heap"

type hpIntMax []int

func (this hpIntMax) Len() int {
	return len(this)
}

func (this hpIntMax) Less(i, j int) bool {
	return this[i] > this[j]
}

func (this hpIntMax) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *hpIntMax) Push(x any) {
	(*this) = append((*this), x.(int))
}

func (this *hpIntMax) Pop() any {
	x := (*this)[len(*this)-1]
	(*this) = (*this)[:len(*this)-1]
	return x
}

type hpIntMin struct {
	hpIntMax
}

func (this hpIntMin) Less(i, j int) bool {
	return this.hpIntMax[i] < this.hpIntMax[j]
}

type MedianFinder struct {
	MaxHp hpIntMax
	MinHp hpIntMin
	Size  int
}

func Constructor295() MedianFinder {
	return MedianFinder{
		Size:  0,
		MaxHp: make([]int, 0, 10),
		MinHp: hpIntMin{
			hpIntMax: make([]int, 0, 10),
		},
	}

}

func (this *MedianFinder) AddNum(num int) {
	l := len(this.MaxHp) + len(this.MinHp.hpIntMax)
	if l == 0 {
		heap.Push(&this.MinHp, num)
		return
	}

	if num >= this.MinHp.hpIntMax[0] {
		heap.Push(&this.MinHp, num)
	} else {
		heap.Push(&this.MaxHp, num)
	}

	for len(this.MinHp.hpIntMax) > len(this.MaxHp)+1 {
		tmp := heap.Pop(&this.MinHp)

		heap.Push(&this.MaxHp, tmp)
	}

	for len(this.MinHp.hpIntMax) < len(this.MaxHp) {
		tmp := heap.Pop(&this.MaxHp)

		heap.Push(&this.MinHp, tmp)
	}

	//fmt.Println(this.MaxHp, this.MinHp, num)

}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.MinHp.hpIntMax) > len(this.MaxHp) {
		return float64(this.MinHp.hpIntMax[0])
	}

	return (float64(this.MinHp.hpIntMax[0]) + float64(this.MaxHp[0])) / 2.0
}
