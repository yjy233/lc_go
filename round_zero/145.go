package roundzero

import "slices"

func postorderTraversal(root *TreeNode) []int {

	st := make([]*TreeNode, 0, 64)

	res := make([]int, 0, 64)

	for root != nil || len(st) > 0 {
		for root != nil {
			st = append(st, root.Left)

			res = append(res, root.Val)
			root = root.Right
		}

		root = st[len(st)-1]
		st = st[:len(st)-1]
	}
	slices.Reverse(res)
	return res
}
