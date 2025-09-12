package roundzero

func partition(head *ListNode, x int) *ListNode {
	moreH := &ListNode{}
	lessH := &ListNode{}

	tmpMoreH := moreH
	tmpLessH := lessH

	for head != nil {
		if head.Val < x {
			tmpLessH.Next = head
			head = head.Next
			tmpLessH = tmpLessH.Next
			tmpLessH.Next = nil
			continue
		}

		tmpMoreH.Next = head
		head = head.Next
		tmpMoreH = tmpMoreH.Next
		tmpMoreH.Next = nil
	}

	if lessH.Next == nil {
		return moreH.Next
	}

	tmpLessH.Next = moreH.Next
	return lessH.Next
}
