package roundzero

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	tmp := &ListNode{
		Next: nil,
	}

	ptr := tmp

	for list1 != nil || list2 != nil {
		if list1 == nil {
			ptr.Next = list2
			break
		}

		if list2 == nil {
			ptr.Next = list1
			break
		}

		if list1.Val < list2.Val {
			tmp.Next = list1
			list1 = list1.Next
			tmp = tmp.Next
			tmp.Next = nil
			continue
		}

		tmp.Next = list2
		list2 = list2.Next
		tmp = tmp.Next
		tmp.Next = nil
		continue

	}
	return ptr.Next
}
