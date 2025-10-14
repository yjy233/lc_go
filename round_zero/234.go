package roundzero

import "fmt"

func isPalindrome234(head *ListNode) bool {
	// a b
	// 1 2 3 4 5
	// 1 2 3 4

	slow := head
	quick := head

	for quick.Next != nil {
		quick = quick.Next

		if quick.Next == nil {
			break
		}
		quick = quick.Next
		slow = slow.Next
	}

	quick = slow.Next
	slow.Next = nil

	quick = reverse234(quick)
	//showList234(quick)
	slow = head
	for slow != nil && quick != nil {
		if slow.Val != quick.Val {
			return false
		}
		slow = slow.Next
		quick = quick.Next
	}

	return true
}

func showList234(head *ListNode) {
	res := ""
	for head != nil {
		res += "->"
		res += fmt.Sprint(head.Val)

		head = head.Next
	}
	fmt.Println(res)
}

func reverse234(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	p0 := head
	p1 := head.Next
	head = p1.Next

	p0.Next = nil

	for head != nil {
		p2 := head.Next

		p1.Next = p0
		p0 = p1
		p1 = head
		p1.Next = nil

		head = p2

	}
	p1.Next = p0
	return p1
}
