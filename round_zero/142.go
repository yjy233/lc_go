package roundzero

func detectCycle(head *ListNode) *ListNode {
	slow, quick := head, head

	hasCycle := false
	for quick != nil && slow != nil {
		quick = quick.Next
		slow = slow.Next

		if quick == nil {
			break
		}

		quick = quick.Next
		if quick == slow {
			hasCycle = true
			break
		}
	}

	//fmt.Println(slow.Val)
	//return nil

	quick = head
	if hasCycle {
		for quick != slow {
			quick = quick.Next
			slow = slow.Next
		}

		return quick
	}

	return nil
}
