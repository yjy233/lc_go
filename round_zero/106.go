package roundzero

func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}

	l := len(postorder)
	val := postorder[l-1]

	indorInd := -1
	for i, v := range inorder {
		if v == val {
			indorInd = i
			break
		}
	}

	var lIn, rIn, lPost, rPost []int

	lIn = inorder[:indorInd]
	if indorInd < l-1 {
		rIn = inorder[indorInd+1:]
	}

	lPost = postorder[:len(lIn)]

	if l > 0 {
		rPost = postorder[len(lPost) : l-1]
	}

	return &TreeNode{
		Val:   val,
		Left:  buildTree106(lIn, lPost),
		Right: buildTree106(rIn, rPost),
	}
}
