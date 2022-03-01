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
		Tail: head,
	}
}

func (ll *LinkedList) AddToEnd(node *Node) {
	// if first element add as head
	if ll.Head == nil {
		ll.Head = node
	}

	node.Prev = ll.Tail
	if ll.Tail != nil {
		ll.Tail.Next = node
	}
	ll.Tail = node
}

func (ll *LinkedList) DeleteFromEnd() *Node {
	if ll.Tail != nil {
		deletedNode := ll.Tail
		if ll.Tail.Prev == nil {
			ll.Tail = nil
		} else {
			ll.Tail = ll.Tail.Prev
			ll.Tail.Next = nil
		}
		return deletedNode
	} else {
		return nil
	}
}

func (ll *LinkedList) Read(index int) *Node {
	node := ll.Head

	for i := 1; i <= index; i++ {
		node = node.Next
		// break loop if requested index is outside list
		if node == nil {
			break
		}
	}

	return node
}
