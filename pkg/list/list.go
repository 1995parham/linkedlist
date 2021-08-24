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
