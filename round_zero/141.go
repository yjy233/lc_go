package roundzero

func hasCycle(head *ListNode) bool {
	slow, quick := head, head

	for quick != nil && slow != nil {
		quick = quick.Next
		slow = slow.Next

		if quick == nil {
			break
		}

		quick = quick.Next
		if quick == slow {
			return true
		}
	}

	return false
}
