package BinaryTree

type Node[T any] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func CreateNode(value any) *Node[any] {
	return &Node[any]{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (node *Node[T]) InOrderTraversal(parent *Node[T]) {
	if parent == nil {
		return
	}
	node.InOrderTraversal(parent.left)
	node.InOrderTraversal(parent.right)
}
