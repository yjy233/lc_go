package roundzero

import "math"

func isBalanced(root *TreeNode) bool {

	res := true
	helper110(root, &res)
	return res
}

func helper110(root *TreeNode, res *bool) int {
	if !(*res) || root == nil {
		return 0
	}

	l := helper110(root.Left, res)
	r := helper110(root.Right, res)

	if math.Abs(float64(l-r)) >= 2 {
		(*res) = false
	}

	return max(l, r) + 1

}
