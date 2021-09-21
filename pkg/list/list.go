package list

import (
	"fmt"
)

type List[T any] struct {
	Head Node[T]
}

func New[T any]() *List[T] {
	return &List[T]{&endNode[T]{}}
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
	l.Head.pushNext(l.Head, nn)
}

func (l *List[T]) Filter(fn func(T) bool) []T {
	r := make([]T, 0)

	for e := l.Head; e.valid(); e = e.next()  {
		data := getData(e)
		if fn(data) {
			r = append(r, data)
		}
	}

	return r
}

func getData[T any](e Node[T]) T {
	if n, ok := e.(*node[T]); ok {
		return n.Data
	}
	return *new(T)
}

func (l *List[T]) String() string {
	r := ""

	for e := l.Head; e.valid(); e = e.next() {
		r += fmt.Sprintf("%v", getData(e))
	}

	return r
}