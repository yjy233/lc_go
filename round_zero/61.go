package roundzero

func rotateRight(head *ListNode, k int) *ListNode {
	l := 0
	tmp := head

	for tmp != nil {
		l++
		tmp = tmp.Next
	}

	if l == 0 {
		return head
	}

	k = k % l
	if k == 0 {
		return head
	}

	res := &ListNode{
		Next: head,
	}

	quick, slow := res, res
	for i := 0; i < k; i++ {
		quick = quick.Next
	}

	for quick != nil && slow != nil {
		quick = quick.Next
		if quick == nil {
			break
		}
		slow = slow.Next
	}

	tmp = slow.Next
	slow.Next = nil
	slow = tmp
	for slow.Next != nil {
		slow = slow.Next
	}
	slow.Next = res.Next
	return tmp

}
