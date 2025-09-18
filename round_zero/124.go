package roundzero

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := root.Val

	helper124(root, &res)
	return res
}

func helper124(root *TreeNode, res *int) int {

	if root == nil {
		return 0
	}

	l := helper124(root.Left, res)
	r := helper124(root.Right, res)

	(*res) = max(root.Val, root.Val+r, root.Val+l, root.Val+r+l, *res)

	return max(root.Val, root.Val+l, root.Val+r)
}
