package roundzero

func removeElements(head *ListNode, val int) *ListNode {
	res := &ListNode{}
	tmp := res

	for head != nil {
		if head.Val == val {
			head = head.Next
			continue
		}

		tmp.Next = head
		head = head.Next
		tmp = tmp.Next
		tmp.Next = nil
	}
	return res.Next
}
