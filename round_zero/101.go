package roundzero

func isSymmetric(root *TreeNode) bool {
	res := true
	if root == nil {
		return res
	}

	helper101(root.Left, root.Right, &res)
	return res
}

func helper101(left *TreeNode, right *TreeNode, res *bool) {
	if !(*res) {
		return
	}

	if left == nil && right == nil {
		return
	}

	if (left == nil && right != nil) ||
		(left != nil && right == nil) {

		(*res) = false
		return
	}

	if left.Val != right.Val {
		*res = false
		return
	}

	helper101(left.Left, right.Right, res)
	helper101(left.Right, right.Left, res)

}
