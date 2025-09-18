package roundzero

import "slices"

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)

	helper107(root, 0, &res)

	slices.Reverse(res)
	return res
}

func helper107(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}

	if level >= len(*res) {
		(*res) = append((*res), []int{})
	}

	helper102(root.Left, level+1, res)
	(*res)[level] = append((*res)[level], root.Val)
	helper102(root.Right, level+1, res)
}
