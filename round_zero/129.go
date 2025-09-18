package roundzero

func sumNumbers(root *TreeNode) int {
	res := 0
	path := make([]int, 0, 100)

	helper129(root, &res, &path)
	return res
}

func helper129(root *TreeNode, res *int, path *[]int) {

	if root == nil {
		return
	}

	(*path) = append((*path), root.Val)
	if root.Left == nil && root.Right == nil {
		tmp := 0
		for _, i := range *path {
			tmp *= 10
			tmp += i
		}
		(*res) += tmp

	} else {
		helper129(root.Left, res, path)
		helper129(root.Right, res, path)
	}
	(*path) = (*path)[:len(*path)-1]
}
