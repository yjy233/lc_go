package roundzero

func reverseKGroup(head *ListNode, k int) *ListNode {

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

	cnt := 1

	termHead := head
	head = head.Next

	for head != nil {
		cnt++
		if cnt%k == 0 {
			//fmt.Println(head.Val)
			termTail := head
			head = head.Next
			termTail.Next = nil

			p0 := termHead
			p1 := termHead.Next
			p0.Next = nil
			for p1 != termTail {
				p2 := p1.Next
				p1.Next = p0

				p0 = p1
				p1 = p2
			}
			p1.Next = p0
			getResTail().Next = p1

			termHead = head
			if head == nil {
				break
			}
			head = head.Next

			cnt = 1
			continue
		}
		head = head.Next
	}

	getResTail().Next = termHead

	return res.Next
}
