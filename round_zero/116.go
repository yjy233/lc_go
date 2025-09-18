package roundzero

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	levelList := make([]*Node, 0)

	helper116(root, 0, &levelList)
	return root
}

func helper116(root *Node, level int, levelList *[]*Node) {

	if root == nil {
		return
	}

	if level < len(*levelList) {
		root.Next = (*levelList)[level]
		(*levelList)[level] = root
	} else {
		(*levelList) = append((*levelList), root)
	}

	helper116(root.Right, level+1, levelList)
	helper116(root.Left, level+1, levelList)

}
