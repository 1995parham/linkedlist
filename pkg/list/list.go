package list

import (
	"fmt"
)

type List[T any] struct {
	Head *Node[T]
}

func New[T any]() *List[T] {
	return &List[T]{
		Head: nil,
	}
}

func (l *List[T]) Len() int {
	e := l.Head
	len := 0

	for e != nil {
		e = e.Next

		len++
	}

	return len
}

func (l *List[T]) PushFront(data T) {
	nn := NewNode(data)
	if l.Head == nil {
		l.Head = nn
	} else {
		nn.Next = l.Head
		l.Head = nn
	}
}

func (l *List[T]) PushBack(data T) {
	nn := NewNode(data)
	if l.Head == nil {
		l.Head = nn
	} else {
		e := l.Head

		for e.Next != nil {
			e = e.Next
		}

		e.Next = nn
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
