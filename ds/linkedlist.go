package ds

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
	}
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList(head *Node) *LinkedList {
	return &LinkedList{
		Head: head,
	}
}
