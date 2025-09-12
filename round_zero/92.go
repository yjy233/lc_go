package roundzero

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	res := &ListNode{}
	res.Next = head

	slow := res
	quick := res
	for i := 0; i < right; i++ {
		quick = quick.Next
	}

	rest := quick.Next
	quick.Next = nil

	for i := 0; i < left-1; i++ {
		slow = slow.Next
	}

	tmp := slow
	slow = slow.Next
	tmp.Next = nil

	tmp = slow.Next
	slow.Next = nil
	for tmp != nil {
		nextTmp := tmp.Next
		tmp.Next = slow
		slow = tmp
		tmp = nextTmp
	}

	tmp = res
	for tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp.Next = slow
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = rest
	return res.Next
}
