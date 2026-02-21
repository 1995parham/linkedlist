package list

import (
	"fmt"
	"iter"
)

// Comparable is a recursive generic constraint (Go 1.26) that requires
// types to be comparable with themselves.
type Comparable[C Comparable[C]] interface {
	Less(C) bool
}

// SortedList is a linked list that maintains elements in ascending order.
// It uses the Go 1.26 recursive generic type feature to constrain T
// to types that can compare with themselves.
type SortedList[T Comparable[T]] struct {
	list *List[T]
}

func NewSorted[T Comparable[T]]() *SortedList[T] {
	return &SortedList[T]{
		list: New[T](),
	}
}

// Insert adds an element in its sorted position (ascending order).
func (sl *SortedList[T]) Insert(data T) {
	nn := newNode(data)

	// Insert at front if list is empty or data comes before head.
	if !sl.list.Head.valid() || data.Less(sl.list.Head.data()) {
		nn.setNext(sl.list.Head)
		sl.list.Head = nn

		return
	}

	// Walk past all elements that are less than data.
	curr := sl.list.Head
	for curr.next().valid() && curr.next().data().Less(data) {
		curr = curr.next()
	}

	nn.setNext(curr.next())
	curr.setNext(nn)
}

func (sl *SortedList[T]) Len() int {
	return sl.list.Len()
}

func (sl *SortedList[T]) Values() iter.Seq[T] {
	return sl.list.Values()
}

func (sl *SortedList[T]) All() iter.Seq2[int, T] {
	return sl.list.All()
}

func (sl *SortedList[T]) Collect() []T {
	return sl.list.Collect()
}

func (sl *SortedList[T]) Filter(fn func(T) bool) iter.Seq[T] {
	return sl.list.Filter(fn)
}

func (sl *SortedList[T]) String() string {
	r := "[ "

	for value := range sl.Values() {
		r += fmt.Sprintf("%v ", value)
	}
	r += "]"

	return r
}
