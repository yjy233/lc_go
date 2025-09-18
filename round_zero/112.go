package roundzero

func hasPathSum(root *TreeNode, targetSum int) bool {

	res := false

	helper112(root, 0, targetSum, &res)
	return res
}

func helper112(root *TreeNode, path, target int, res *bool) {
	if *res || root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if path+root.Val == target {
			*res = true
			return
		}
	}

	helper112(root.Left, path+root.Val, target, res)
	helper112(root.Right, path+root.Val, target, res)
}
