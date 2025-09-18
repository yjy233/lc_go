package roundzero

import (
	"slices"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	helper103(root, 0, &res)

	for i := 1; i < len(res); i++ {
		if i%2 == 1 {
			slices.Reverse(res[i])
		}
	}
	return res
}

func helper103(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}

	if level >= len(*res) {
		(*res) = append((*res), []int{})
	}

	helper103(root.Left, level+1, res)
	(*res)[level] = append((*res)[level], root.Val)
	helper103(root.Right, level+1, res)
}
