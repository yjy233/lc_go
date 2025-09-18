package roundzero

func buildTree(preorder []int, inorder []int) *TreeNode {
	//fmt.Println(preorder, inorder)
	if len(preorder) <= 0 || len(inorder) <= 0 || len(preorder) != len(inorder) {
		return nil
	}

	val := preorder[0]

	inorderInd := -1
	for i, v := range inorder {
		if v == val {
			inorderInd = i
			break
		}
	}

	var lPre, rPre, lIn, rIn []int
	if inorderInd >= 1 {
		lPre = preorder[1 : inorderInd+1]
	}

	if inorderInd < len(preorder) {
		rPre = preorder[inorderInd+1 : len(preorder)]
	}

	if inorderInd > 0 {
		lIn = inorder[:inorderInd]
	}

	if inorderInd < len(inorder)-1 {
		rIn = inorder[inorderInd+1:]
	}

	return &TreeNode{
		Val:   val,
		Left:  buildTree(lPre, lIn),
		Right: buildTree(rPre, rIn),
	}

}
