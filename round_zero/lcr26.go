package roundzero

func reorderList(head *ListNode) {

	// 1. 2 3 4.  1,2.  2,4
	// 1 2. 3 4  5  1,2  2,4  3,5
	res := &ListNode{
		Next: head,
	}

	quick := res
	slow := res

	for quick.Next != nil {
		quick = quick.Next
		slow = slow.Next

		if quick.Next == nil {
			break
		}

		quick = quick.Next
	}

	quick = slow.Next
	slow.Next = nil

	quick = r26(quick)
	slow = res.Next
	res.Next = nil
	tmp := res

	for slow != nil || quick != nil {
		if slow != nil {
			tmp.Next = slow
			slow = slow.Next
			tmp = tmp.Next
			tmp.Next = nil
		}

		if quick != nil {
			tmp.Next = quick
			quick = quick.Next
			tmp = tmp.Next
			tmp.Next = nil
		}

	}
}

func r26(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p0 := head
	p1 := head.Next

	p0.Next = nil

	for p1 != nil {
		p2 := p1.Next

		p1.Next = p0
		p0 = p1
		p1 = p2
	}

	return p0
}
