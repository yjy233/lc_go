package roundzero

import "container/heap"

type heapL []*ListNode

func (h *heapL) Len() int {
	return len(*h)
}

func (h *heapL) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *heapL) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heapL) Push(x any) {
	(*h) = append((*h), x.(*ListNode))
}

func (h *heapL) Pop() any {
	if len(*h) <= 0 {
		return nil
	}

	v := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return v
}

func mergeKListsV2(lists []*ListNode) *ListNode {
	hp := &heapL{}

	for _, ptr := range lists {
		if ptr == nil {
			continue
		}

		(*hp) = append((*hp), ptr)
	}

	heap.Init(hp)

	res := new(ListNode)
	tmp := res

	for hp.Len() > 0 {
		ptr := heap.Pop(hp)
		node, ok := ptr.(*ListNode)
		if node != nil && ok {
			tmp.Next = node
			node = node.Next

			tmp = tmp.Next
			tmp.Next = nil
		}

		if node != nil {
			heap.Push(hp, node)
		}

	}

	return res.Next
}
