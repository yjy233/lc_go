package roundzero

func connectV3(root *Node) *Node {
	levelList := make([]*Node, 0)

	helper116(root, 0, &levelList)
	return root
}

func helper117(root *Node, level int, levelList *[]*Node) {

	if root == nil {
		return
	}

	if level < len(*levelList) {
		root.Next = (*levelList)[level]
		(*levelList)[level] = root
	} else {
		(*levelList) = append((*levelList), root)
	}

	helper117(root.Right, level+1, levelList)
	helper117(root.Left, level+1, levelList)

}
