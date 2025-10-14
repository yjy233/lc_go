package roundzero

func removeNthFromEndV3(head *ListNode, n int) *ListNode {

	res := &ListNode{
		Next: head,
	}

	quick := res
	slow := res
	for i := 0; i < n; i++ {
		quick = quick.Next
	}
	for quick.Next != nil {
		quick = quick.Next
		slow = slow.Next
	}

	if slow.Next != nil {
		slow.Next = slow.Next.Next
	} else {
		slow.Next = nil
	}

	return res.Next
}
