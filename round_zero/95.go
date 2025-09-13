package roundzero

func generateTrees(n int) []*TreeNode {
	return helper95(1, n)
}

func helper95(st, ed int) []*TreeNode {
	if st > ed || st <= 0 {
		return []*TreeNode{
			nil,
		}
	}

	res := make([]*TreeNode, 0, (ed - st + 1))
	for i := st; i <= ed; i++ {
		lc := helper95(st, i-1)
		rc := helper95(i+1, ed)

		for li := 0; li < len(lc); li++ {
			for ri := 0; ri < len(rc); ri++ {
				res = append(res, genNewNode(i, lc[li], rc[ri]))
			}
		}
	}
	return res
}

func genNewNode(n int, l, r *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   n,
		Left:  l,
		Right: r,
	}
}
