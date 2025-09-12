package roundzero

func isValidBST(root *TreeNode) bool {
	res := true
	helper98(root, &res)
	return res
}

func helper98(root *TreeNode, res *bool) []int {
	if !(*res) || root == nil {
		return nil
	}

	left := helper98(root.Left, res)
	if (*res) == false {
		return nil
	}

	if left != nil {
		if left[1] >= root.Val {
			(*res) = false
			return nil
		}
	}

	right := helper98(root.Right, res)
	if (*res) == false {
		return nil
	}

	if right != nil {
		if right[0] <= root.Val {
			(*res) = false
			return nil
		}
	}

	ret := []int{root.Val, root.Val}
	if left != nil {
		ret[0] = left[0]
	}
	if right != nil {
		ret[1] = right[1]
	}

	return ret
}
