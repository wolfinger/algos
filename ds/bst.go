package ds

type BSTNode struct {
	Data  int
	Left  *BSTNode
	Right *BSTNode
}

func NewBSTNode(data int) *BSTNode {
	return &BSTNode{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func BSTSearch(node *BSTNode, value int) *BSTNode {
	if (node == nil) || (node.Data == value) {
		return node
	} else if value < node.Data {
		return BSTSearch(node.Left, value)
	} else {
		return BSTSearch(node.Right, value)
	}
}
