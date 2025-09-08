package roundzero

import (
	"container/heap"
	"fmt"
)

/*
本题关键，就是 pop push 要指针 * 一定要注意
*/
type hp []*ListNode

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *hp) Pop() any {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{}
	ptr := res

	h := hp{}
	for _, item := range lists {
		if item != nil {
			h = append(h, item)
		}
	}
	fmt.Println("hjiwfwefwef we")
	heap.Init(&h)
	fmt.Println("-----------------")
	for len(h) > 0 {
		ln := heap.Pop(&h)
		lnp, ok := ln.(*ListNode)

		if ok && lnp != nil {
			fmt.Println(lnp.Val)
			ptr.Next = lnp
			lnp = lnp.Next
			ptr = ptr.Next
			ptr.Next = nil

			if lnp != nil {
				fmt.Println(lnp.Val)
				heap.Push(&h, lnp)
			}
		}
	}

	return res.Next
}
