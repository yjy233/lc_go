package roundzero

import "container/heap"

type hp []*ListNode

func (this hp) Len() int {
	return len(this)
}

func (this hp) Less(i, j int) bool {
	return (this)[i].Val < (this)[j].Val
}

func (this hp) Swap(i, j int) {
	(this)[i], (this)[j] = (this)[j], (this)[i]
}

func (this *hp) Push(x any) {
	(*this) = append((*this), x.(*ListNode))
}

func (this *hp) Pop() any {
	v := (*this)[len(*this)-1]
	(*this) = (*this)[:len(*this)-1]
	return v
}

func mergeKLists(lists []*ListNode) *ListNode {

	newList := &hp{}
	for _, i := range lists {
		if i == nil {
			continue
		}

		(*newList) = append((*newList), i)
	}
	heap.Init(newList)
	res := &ListNode{}
	tmp := res

	for len(*newList) > 0 {
		ptr := heap.Pop(newList)

		if ptr == nil {
			continue
		}

		ptrNode := ptr.(*ListNode)
		tmp.Next = ptrNode
		ptrNode = ptrNode.Next
		tmp = tmp.Next
		tmp.Next = nil

		if ptrNode != nil {
			heap.Push(newList, ptrNode)
		}
	}
	return res.Next
}
