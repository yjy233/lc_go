package roundzero

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res ListNode = ListNode{}
	tmp := &res
	fmt.Println("xxx")
	a := append([]int{}, 3)
	fmt.Println(a)

	jin := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = l2
			l2 = nil
			continue
		}

		l1.Val += jin
		if l2 != nil {
			l1.Val += l2.Val
			l2 = l2.Next
		}

		jin = l1.Val / 10
		l1.Val = l1.Val % 10

		tmp.Next = l1
		l1 = l1.Next
		tmp = tmp.Next
		tmp.Next = nil

	}

	if jin > 0 {
		tmp.Next = new(ListNode)
		tmp.Next.Val = 1
		tmp.Next.Next = nil
	}

	return res.Next
}
