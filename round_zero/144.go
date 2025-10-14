package roundzero

func preorderTraversal(root *TreeNode) []int {

	st := make([]*TreeNode, 0, 64)

	res := make([]int, 0, 64)

	for root != nil || len(st) > 0 {

		for root != nil {
			st = append(st, root.Right)
			res = append(res, root.Val)
			root = root.Left
		}

		root = st[len(st)-1]
		st = st[:len(st)-1]

	}
	return res
}
