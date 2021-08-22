package node

type Node[T any] struct {
	Data T
	Next *Node[T]
}


func New[T any](data T) *Node[T] {
	return &Node[T]{
		Data: data,
		Next: nil,
	}
}
