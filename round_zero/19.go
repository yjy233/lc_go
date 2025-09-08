package roundzero

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 0->1->2->3->4->5 2

	tmp := &ListNode{}
	tmp.Next = head

	quick := tmp
	for i := 0; i < n; i++ {
		quick = quick.Next
	}

	slow := tmp
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next
	}

	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}

	return tmp.Next
}
