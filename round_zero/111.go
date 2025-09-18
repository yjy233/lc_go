package roundzero

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := -1
	helper111(root, &res, 1)
	return res
}

func helper111(root *TreeNode, res *int, dep int) {

	if root == nil || (*res != -1 && dep >= *res) {
		return
	}

	if root.Left == nil && root.Right == nil {
		if (*res == -1) || dep < *res {
			(*res) = dep
		}
		return
	}

	helper111(root.Left, res, dep+1)
	helper111(root.Right, res, dep+1)
}
