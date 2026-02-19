package list

import (
	"fmt"
	"iter"
	"slices"
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
	for range l.Values() {
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

func (l *List[T]) Filter(fn func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range l.Values() {
			if fn(value) {
				if !yield(value) {
					return
				}
			}
		}
	}
}

// Values returns an iterator over all element values.
// This follows the slices.Values convention.
func (l *List[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := l.Head; e.valid(); e = e.next() {
			if !yield(e.data()) {
				return
			}
		}
	}
}

// All returns an iterator over index-value pairs.
// This follows the slices.All convention (Go 1.23+).
func (l *List[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		index := 0
		for e := l.Head; e.valid(); e = e.next() {
			if !yield(index, e.data()) {
				return
			}
			index++
		}
	}
}

// Collect materializes the list into a slice using slices.Collect (Go 1.23+).
func (l *List[T]) Collect() []T {
	return slices.Collect(l.Values())
}

func (l *List[T]) String() string {
	r := "[ "

	for value := range l.Values() {
		r += fmt.Sprintf("%v ", value)
	}
	r += "]"

	return r
}

// Map transforms each element of an iterator using the given function,
// returning a new iterator. This is a standalone function because Go methods
// cannot introduce additional type parameters (Go generics constraint).
func Map[T, U any](s iter.Seq[T], fn func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range s {
			if !yield(fn(v)) {
				return
			}
		}
	}
}
