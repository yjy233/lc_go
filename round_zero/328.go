package roundzero

func oddEvenList(head *ListNode) *ListNode {

	single := &ListNode{}
	s := single
	double := &ListNode{}
	d := double

	for head != nil {
		s.Next = head
		head = head.Next
		s = s.Next
		s.Next = nil

		if head == nil {
			break
		}

		d.Next = head
		head = head.Next
		d = d.Next
		d.Next = nil
	}

	//fmt.Println("end")
	s.Next = double.Next
	return single.Next
}
