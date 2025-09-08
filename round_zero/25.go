package roundzero

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	res := &ListNode{}
	getPtrF := func() *ListNode {
		ptr := res
		for ptr.Next != nil {
			ptr = ptr.Next
		}
		return ptr
	}

	cnt := 1
	termHead := head
	termTail := head
	head = head.Next

	for head != nil {
		cnt++

		if cnt%k == 0 {
			termTail = head
			nextHead := head.Next

			termTail.Next = nil
			tmpPtr := termHead.Next
			termHead.Next = nil
			for tmpPtr != termTail {
				nextT := tmpPtr.Next
				tmpPtr.Next = termHead

				termHead = tmpPtr
				tmpPtr = nextT
			}
			termTail.Next = termHead

			getPtrF().Next = termTail

			cnt = 1
			head = nextHead
			termHead = head
			termTail = head
			if head == nil {
				break
			}
		}

		head = head.Next
	}

	getPtrF().Next = termHead

	return res.Next
}
