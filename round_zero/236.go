package roundzero

func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {

	res := &TreeNode{}
	helper236(root, p, q, &res)
	return res
}

func helper236(root, p, q *TreeNode, res **TreeNode) int {
	if root == nil {
		return -1
	}

	l := helper236(root.Left, p, q, res)
	r := helper236(root.Right, p, q, res)
	if l > 0 && r > 0 {
		(*res) = root
		return 3
	}

	if root == p {
		if l == 2 || r == 2 {
			(*res) = root
			return 3
		}
		return 1
	} else if root == q {
		if l == 1 || r == 1 {
			(*res) = root
			return 3
		}
		return 2
	}

	if l > 0 {
		return l
	}
	if r > 0 {
		return r
	}
	return -1
}
