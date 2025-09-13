package roundzero

func recoverTree(root *TreeNode) {
	nodeList := make([]*TreeNode, 0, 100)
	helper99(root, &nodeList)

	firstNode := (*TreeNode)(nil)
	secondNode := (*TreeNode)(nil)

	for i := 1; i < len(nodeList); i++ {
		if nodeList[i].Val < nodeList[i-1].Val {
			if firstNode == nil {
				firstNode = nodeList[i-1]
				secondNode = nodeList[i]
			} else {
				secondNode = nodeList[i]
			}
		}
	}

	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
}

func helper99(root *TreeNode, nodeList *[]*TreeNode) {
	if root == nil {
		return
	}

	helper99(root.Left, nodeList)
	(*nodeList) = append((*nodeList), root)
	helper99(root.Right, nodeList)
}
