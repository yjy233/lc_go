package roundzero

import "container/heap"

type hp23 []*ListNode

func (t hp23) Len() int {
	return len(t)
}

func (t hp23) Less(i, j int) bool {
	return t[i].Val < t[j].Val
}

func (t hp23) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (this *hp23) Push(x any) {
	(*this) = append((*this), x.(*ListNode))
}

func (this *hp23) Pop() any {
	x := (*this)[len(*this)]
	(*this) = (*this)[:len(*this)-1]
	return x
}

func mergeKListsV3(lists []*ListNode) *ListNode {

	hp := (hp23)(make([]*ListNode, 0))
	for i := range lists {
		if lists[i] == nil {
			continue
		}

		hp = append(hp, lists[i])
	}

	res := &ListNode{}
	tmp := res

	heap.Init(&hp)
	for len(hp) > 0 {
		node := heap.Pop(&hp).(*ListNode)

		if node == nil {
			continue
		}

		tmp.Next = node
		node = node.Next
		tmp = tmp.Next
		tmp.Next = nil

		if node != nil {
			heap.Push(&hp, node)
		}

	}

	return res.Next
}
