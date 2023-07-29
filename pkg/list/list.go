package list

import (
	"fmt"
)

type List[T any] struct {
	Head Node[T]
}

func New[T any]() *List[T] {
	return &List[T]{
		Head: newEndNode[T](),
	}
}

func (l *List[T]) setNext(node Node[T]) {
	l.Head = node
}

func (l *List[T]) Len() int {
	length := 0
	for e := l.Head; e.valid(); e = e.next() {
		length++
	}

	return length
}

func (l *List[T]) PushFront(data T) {
	nn := newNode(data)
	nn.setNext(l.Head)
	l.Head = nn
}

func (l *List[T]) PushBack(data T) {
	nn := newNode(data)

	l.Head.pushNext(l, nn)
}

func (l *List[T]) Filter(fn func(T) bool) []T {
	r := make([]T, 0)

	for e := l.Head; e.valid(); e = e.next() {
		data := e.data()
		if fn(data) {
			r = append(r, data)
		}
	}

	return r
}

func (l *List[T]) String() string {
	r := "[ "

	for e := l.Head; e.valid(); e = e.next() {
		r += fmt.Sprintf("%v ", e.data())
	}
	r += "]"

	return r
}
