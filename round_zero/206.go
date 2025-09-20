package roundzero

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	first := head
	head = head.Next
	if head == nil {
		return first
	}

	first.Next = nil
	second := head
	head = head.Next
	second.Next = nil

	for head != nil {
		second.Next = first

		first = second
		second = head
		head = head.Next
	}

	second.Next = first
	return second
}
