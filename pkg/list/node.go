package list

type Linker[T any] interface {
	setNext(Node[T])
}

// Node interface defines an interface for nodes in the linked list
// so we can iterate them easier.
type Node[T any] interface {
	Linker[T]
	pushNext(last Linker[T], new Node[T])
	valid() bool
	next() Node[T]
}

type node[T any] struct {
	Data T
	Next Node[T]
}

func newNode[T any](data T) *node[T] {
	return &node[T]{
		Data: data,
		Next: nil,
	}
}

func (n *node[T]) pushNext(_ Linker[T], new Node[T]) {
	n.Next.pushNext(n, new)
}

func (n *node[T]) setNext(node Node[T]) {
	n.Next = node
}

func (n *node[T]) valid() bool {
	return true
}

func (n *node[T]) next() Node[T] {
	return n.Next
}

// endNode indicates end of the linked list and it will reduce our checks
// by automatically handles the edge conditions.
type endNode[T any] struct{}

func newEndNode[T any]() *endNode[T] {
	return &endNode[T]{}
}

func (e *endNode[T]) valid() bool {
	return false
}

func (e *endNode[T]) next() Node[T] {
	return e
}

func (e *endNode[T]) setNext(Node[T]) {}

func (e *endNode[T]) pushNext(previous Linker[T], new Node[T]) {
	previous.setNext(new)
	new.setNext(e)
}
