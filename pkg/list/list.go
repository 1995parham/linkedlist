package list

import (
	"fmt"

	"github.com/1995parham/linkedlist/pkg/node"
)

type List[T any] struct {
	Head *node.Node[T]
}

func New[T any]() *List[T] {
	return &List[T]{
		Head: nil,
	}
}

func (l *List[T]) PushFront(data T) {
	nn := node.New(data)
	if l.Head == nil {
		l.Head = nn
	} else {
		nn.Next = l.Head
		l.Head = nn
	}
}

func (l *List[T]) Filter(fn func (T) bool) []T {
	e := l.Head
	r := make([]T, 0)

	for e != nil {
		if fn(e.Data) {
			r = append(r, e.Data)
		}
		e = e.Next
	}

	return r
}

func (l *List[T]) String() string {
	e := l.Head
	r := ""

	for e != nil {
		r += fmt.Sprintf("%v", e.Data)
		e = e.Next

		if e != nil {
			r += " "
		}
	}

	return r
}
