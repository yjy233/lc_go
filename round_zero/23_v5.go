package roundzero

import "container/heap"

type hpV5 []*ListNode

func (this hpV5) Len() int {
	return len(this)
}

func (this hpV5) Less(i, j int) bool {
	return (this)[i].Val < (this)[j].Val
}

func (this hpV5) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *hpV5) Pop() any {
	v := (*this)[len(*this)-1]
	(*this) = (*this)[:len(*this)-1]
	return v
}

func (this *hpV5) Push(x any) {
	(*this) = append((*this), x.(*ListNode))
}

func mergeKListsV5(lists []*ListNode) *ListNode {
	h := make([]*ListNode, 0, len(lists))
	hp := (hpV5)(h)

	for _, ptr := range lists {
		if ptr == nil {
			continue
		}

		hp = append(hp, ptr)
	}

	heap.Init(&hp)

	res := &ListNode{}
	tmp := res

	for len(hp) > 0 {
		ptr := heap.Pop(&hp).(*ListNode)

		tmp.Next = ptr
		ptr = ptr.Next
		tmp = tmp.Next
		tmp.Next = nil

		if ptr != nil {
			heap.Push(&hp, ptr)
		}
	}

	return res.Next
}
