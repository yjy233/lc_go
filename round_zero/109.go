package roundzero

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	tmp := head
	l := 0
	for head != nil {
		l++
		head = head.Next
	}

	fackHead := &ListNode{
		Next: tmp,
	}

	tmp = fackHead
	for i := 0; i < l/2; i++ {
		tmp = tmp.Next
	}

	val := tmp.Next.Val
	head = tmp.Next.Next

	tmp.Next = nil
	return &TreeNode{
		Val:   val,
		Left:  sortedListToBST(fackHead.Next),
		Right: sortedListToBST(head),
	}
}
