package roundzero

func deleteDuplicates83(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	res := &ListNode{
		Next: head,
	}

	tmp := head
	head = head.Next
	tmp.Next = nil

	pre := tmp.Val
	for head != nil {
		if pre == head.Val {
			head = head.Next
			continue
		}

		tmp.Next = head
		head = head.Next
		tmp = tmp.Next
		tmp.Next = nil

		pre = tmp.Val
	}
	return res.Next
}
