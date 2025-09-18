package roundzero

func maxDepth(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}

	helper104(root, &res, 1)
	return res
}

func helper104(root *TreeNode, res *int, level int) {
	if root == nil {
		return
	}

	(*res) = max(*res, level)

	helper104(root.Left, res, level+1)
	helper104(root.Right, res, level+1)
}
