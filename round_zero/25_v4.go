package roundzero

func reverseKGroupV4(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}

	res := &ListNode{}

	getResTail := func() *ListNode {
		tmp := res
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		return tmp
	}

	termHead := head
	head = head.Next
	cnt := 1
	for head != nil {
		cnt++
		if cnt%k == 0 {
			termTail := head
			nextHead := head.Next
			p1 := termHead.Next
			termHead.Next = nil

			for p1 != termTail {
				p2 := p1.Next

				p1.Next = termHead
				termHead = p1
				p1 = p2
			}

			termTail.Next = termHead
			getResTail().Next = termTail

			termHead = nextHead
			if termHead == nil {
				break
			}
			head = termHead.Next
			cnt = 1
			continue
		}

		head = head.Next

	}

	getResTail().Next = termHead

	return res.Next
}
