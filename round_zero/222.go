package roundzero

func countNodes(root *TreeNode) int {
	res := 0

	helper222(root, &res)
	return res
}

func helper222(root *TreeNode, res *int) {
	if root == nil {
		return
	}

	(*res)++

	helper222(root.Left, res)
	helper222(root.Right, res)
}
