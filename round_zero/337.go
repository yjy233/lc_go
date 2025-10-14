package roundzero

func robV3(root *TreeNode) int {

	r0, r1 := helper337(root)
	return max(r0, r1, 0)
}

func helper337(root *TreeNode) (int, int) {
	if root == nil {
		return -1, -1
	}

	v := root.Val

	l0, l1 := helper337(root.Left)
	r0, r1 := helper337(root.Right)

	return max(0, l1+r1, l1, r1) + v, max(l0+r0, r0, l0, 0)
}
