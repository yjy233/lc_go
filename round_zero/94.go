package roundzero

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	helper94(root, &res)
	return res
}

func helper94(root *TreeNode, res *[]int) {

	if root == nil {
		return
	}

	helper94(root.Left, res)
	(*res) = append((*res), root.Val)
	helper94(root.Right, res)
}
