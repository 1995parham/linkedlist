package list

type Node[T any] struct {
	Data T
	Next *Node[T]
}


func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		Data: data,
		Next: nil,
	}
}
