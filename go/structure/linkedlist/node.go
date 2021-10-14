package linkedlist

type Node struct {
	Val  interface{}
	Prev *Node
	Next *Node
}

func NewNode(val interface{}) *Node {
	return &Node{val, nil, nil}
}
