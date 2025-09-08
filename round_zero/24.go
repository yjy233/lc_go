package roundzero

func swapPairs(head *ListNode) *ListNode {
	res := new(ListNode)
	tmp := res

	for head != nil {
		p0 := head
		p1 := p0.Next
		if p1 == nil {
			tmp.Next = p0
			break
		}

		head = p1.Next
		p0.Next = nil
		p1.Next = nil

		tmp.Next = p1
		tmp = tmp.Next
		tmp.Next = p0
		tmp = tmp.Next
	}

	return res.Next
}
