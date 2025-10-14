package roundzero

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode = &TreeNode{}

	helper235(root, p, q, &res)
	return res
}

func helper235(root, p, q *TreeNode, res **TreeNode) int {
	if root == nil {
		return -1
	}

	l := helper235(root.Left, p, q, res)
	r := helper235(root.Right, p, q, res)
	//fmt.Println(root.Val, l, r)
	if l > 0 && r > 0 {
		(*res) = root
	} else if root == q {
		//fmt.Println("fwfewfwe ",root.Val, l, r)
		if l == 2 || r == 2 {
			(*res) = root
			return 1
		}
		return 1
	} else if root == p {
		//fmt.Println("mwoifowenfe ", root.Val)
		if l == 1 || r == 1 {
			(*res) = root
			return 2
		}
		return 2
	}

	if r > 0 {
		return r
	}
	if l > 0 {
		return l
	}
	return -1
}
