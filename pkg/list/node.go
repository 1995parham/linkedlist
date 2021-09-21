package list

type Node[T any] interface {
	pushNext(last, new Node[T])
	valid() bool
	next() Node[T]
	setNext(Node[T])
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

func (n *node[T]) pushNext(_, new Node[T]) {
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

func (e *endNode[T]) setNext(Node[T]) {
	return
}

func (e *endNode[T]) pushNext(previous, new Node[T]) {
	previous.setNext(new)
	new.setNext(e)
}
