package roundzero

func deleteDuplicates(head *ListNode) *ListNode {

	res := &ListNode{
		Next: nil,
	}

	if head == nil {
		return head
	}

	tmp := res

	pre := head.Val
	slow := head.Next
	if slow == nil {
		return head
	}

	quick := slow.Next
	if quick == nil {
		if pre == slow.Val {
			return nil
		}
		return head
	}

	if pre != slow.Val {
		tmp.Next = head
		tmp = head
		tmp.Next = nil
	}

	for quick != nil {
		if pre != slow.Val && slow.Val != quick.Val {
			tmp.Next = slow

			slow = quick
			quick = quick.Next
			tmp = tmp.Next

			tmp.Next = nil
			pre = tmp.Val
			continue
		}

		pre = slow.Val
		slow = quick
		quick = quick.Next
	}

	if pre != slow.Val {
		tmp.Next = slow
		slow.Next = nil
	}

	return res.Next
}
