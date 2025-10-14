package roundzero

func kthSmallest(root *TreeNode, k int) int {
	var res *int = nil
	turn := 0
	helper230(root, k, &turn, &res)
	if res == nil {
		return 0
	}
	return (*res)
}

func helper230(root *TreeNode, k int, turn *int, res **int) {
	if root == nil {
		return
	}

	helper230(root.Left, k, turn, res)

	(*turn)++
	if (*turn) == k {
		(*res) = &root.Val
	}

	helper230(root.Right, k, turn, res)
}
