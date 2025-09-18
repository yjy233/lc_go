package roundzero

import "math"

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	res := &ListNode{
		Val:  int(math.MinInt32),
		Next: head,
	}

	head = head.Next
	res.Next.Next = nil

	for head != nil {
		now := head
		head = head.Next
		now.Next = nil

		pre, next := res, res.Next

		for pre.Val <= now.Val && (next != nil && next.Val <= now.Val) {
			pre = pre.Next
			next = next.Next
		}
		//fmt.Println(pre.Val)
		pre.Next = now
		now.Next = next
	}

	return res.Next
}
