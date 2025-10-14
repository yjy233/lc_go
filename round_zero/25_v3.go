package roundzero

func reverseKGroupV3(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil || head.Next == nil {
		return head
	}

	res := new(ListNode)
	getResF := func() *ListNode {
		head := res
		for head.Next != nil {
			head = head.Next
		}
		return head
	}

	cnt := 1
	termHead := head
	head = head.Next

	for head != nil {
		cnt++
		if cnt%k == 0 {
			nextHead := head.Next

			termTail := head
			termTail.Next = nil
			p1 := termHead.Next
			termHead.Next = nil

			for p1 != termTail {
				p2 := p1.Next

				p1.Next = termHead
				termHead = p1
				p1 = p2
			}

			p1.Next = termHead

			getResF().Next = termTail
			cnt = 1
			termHead = nextHead
			if termHead == nil {
				break
			}
			head = termHead.Next
			continue
		}
		head = head.Next
	}

	getResF().Next = termHead

	return res.Next
}
