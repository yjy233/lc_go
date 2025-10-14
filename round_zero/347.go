package roundzero

import "container/heap"

var (
	n2Cnt = make(map[int]int)
)

type hp347 []int

func (this hp347) Len() int {
	return len(this)
}

func (this hp347) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this hp347) Less(i, j int) bool {
	//fmt.Println("--- ",i,j, n2Cnt[i], n2Cnt[j])
	return n2Cnt[this[i]] > n2Cnt[this[j]]
}

func (this *hp347) Pop() any {
	x := (*this)[len(*this)-1]

	(*this) = (*this)[:len(*this)-1]
	return x
}

func (this *hp347) Push(x any) {
	(*this) = append((*this), x.(int))
}

func topKFrequent(nums []int, k int) []int {
	n2Cnt = make(map[int]int)
	for _, n := range nums {
		n2Cnt[n] += 1
	}

	tmp := make([]int, 0, len(nums))
	for k := range n2Cnt {
		tmp = append(tmp, k)
	}

	tmpPtr := (hp347)(tmp)
	heap.Init(&tmpPtr)
	//fmt.Println(tmp)
	res := make([]int, 0, k)

	for i := 0; i < k; i++ {
		h := heap.Pop(&tmpPtr)
		res = append(res, h.(int))
		//fmt.Println(tmp)
	}
	return res
}
