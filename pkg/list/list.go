// Package list implements a generic singly linked list, plus a sorted
// variant built on top of it. It exists to exercise Go generics — type
// parameters, recursive constraints, range-over-func iterators, and
// generic methods — on a data structure that used to need code generation.
package list

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

// List is a singly linked list of T. The zero value is not usable; build one
// with New, which seeds Head with the end-of-list sentinel.
type List[T any] struct {
	Head Node[T]
}

// New returns an empty list ready for use.
func New[T any]() *List[T] {
	return &List[T]{
		Head: newEndNode[T](),
	}
}

func (l *List[T]) setNext(node Node[T]) {
	l.Head = node
}

// Len reports the number of elements in the list. It walks the list, so it
// runs in O(n).
func (l *List[T]) Len() int {
	length := 0
	for range l.Values() {
		length++
	}

	return length
}

// PushFront inserts data at the head of the list in O(1).
func (l *List[T]) PushFront(data T) {
	nn := newNode(data)
	nn.setNext(l.Head)
	l.Head = nn
}

// PushBack appends data to the tail of the list. It walks to the end, so it
// runs in O(n).
func (l *List[T]) PushBack(data T) {
	nn := newNode(data)

	l.Head.pushNext(l, nn)
}

// Filter returns an iterator over the elements for which fn reports true.
// The list is walked lazily as the iterator is consumed.
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

// String renders the list as a bracketed, space-separated sequence.
func (l *List[T]) String() string {
	var sb strings.Builder

	sb.WriteString("[ ")

	for value := range l.Values() {
		fmt.Fprintf(&sb, "%v ", value)
	}

	sb.WriteString("]")

	return sb.String()
}

// Map transforms each element of the list using the given function, returning a
// new iterator over the results.
//
// This is a generic method: it introduces its own type parameter U on top of the
// receiver's T. Before Go 1.27 methods could not declare type parameters, so this
// had to be written as a standalone function taking an iter.Seq[T]. Go 1.27 lifts
// that restriction (proposal golang.org/issue/49085), letting Map read the
// receiver directly while still changing the element type from T to U.
func (l *List[T]) Map[U any](fn func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for value := range l.Values() {
			if !yield(fn(value)) {
				return
			}
		}
	}
}
