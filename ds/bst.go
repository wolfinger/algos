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

func BSTInsert(node *BSTNode, value int) *BSTNode {
	if value < node.Data {
		if node.Left == nil {
			node.Left = NewBSTNode(value)
			return node.Left
		} else {
			return BSTInsert(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = NewBSTNode(value)
			return node.Right
		} else {
			return BSTInsert(node.Right, value)
		}
	}
}
