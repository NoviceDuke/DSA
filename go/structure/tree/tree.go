package tree

type Node struct {
	val   int
	left  *Node
	right *Node
}

func NewNode(val int) *Node {
	return &Node{val, nil, nil}
}

type BSTree struct {
	Root *Node
}

// calDepth helper function  for tree depth
func calDepth(n *Node, depth int) int {
	if n == nil {
		return depth
	}
	return Max(calDepth(n.left, depth+1), calDepth(n.right, depth+1))
}
func Insert(root *Node, val int) *Node {
	if root == nil {
		return NewNode(val)
	}
	if val < root.val {
		root.left = Insert(root.left, val)
	} else {
		root.right = Insert(root.right, val)
	}
	return root
}
func (t *BSTree) Depth() int {
	return calDepth(t.Root, 0)
}
func InOrderSuccessor(root *Node) *Node {
	cur := root
	if cur.left != nil {
		cur = cur.left
	}
	return cur
}
func inOrderRecursive(n *Node, traversal *[]int) {
	if n != nil {
		inOrderRecursive(n.left, traversal)
		*traversal = append(*traversal, n.val)
		inOrderRecursive(n.right, traversal)
	}
}
func InOrder(root *Node) []int {
	traversal := make([]int, 0)

	inOrderRecursive(root, &traversal)

	return traversal
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
