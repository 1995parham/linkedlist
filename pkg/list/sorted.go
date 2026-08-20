package list

import (
	"fmt"
	"iter"
	"strings"
)

// Comparable is a recursive generic constraint (Go 1.26) that requires
// types to be comparable with themselves.
type Comparable[C Comparable[C]] interface {
	Less(other C) bool
}

// SortedList is a linked list that maintains elements in ascending order.
// It uses the Go 1.26 recursive generic type feature to constrain T
// to types that can compare with themselves.
type SortedList[T Comparable[T]] struct {
	list *List[T]
}

// NewSorted returns an empty sorted list ready for use.
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

// Len reports the number of elements in the list.
func (sl *SortedList[T]) Len() int {
	return sl.list.Len()
}

// Values returns an iterator over the elements in ascending order.
func (sl *SortedList[T]) Values() iter.Seq[T] {
	return sl.list.Values()
}

// All returns an iterator over index-value pairs in ascending order.
func (sl *SortedList[T]) All() iter.Seq2[int, T] {
	return sl.list.All()
}

// Collect materializes the list into a sorted slice.
func (sl *SortedList[T]) Collect() []T {
	return sl.list.Collect()
}

// Filter returns an iterator over the elements for which fn reports true,
// preserving the sorted order.
func (sl *SortedList[T]) Filter(fn func(T) bool) iter.Seq[T] {
	return sl.list.Filter(fn)
}

// String renders the list as a bracketed, space-separated sequence.
func (sl *SortedList[T]) String() string {
	var sb strings.Builder

	sb.WriteString("[ ")

	for value := range sl.Values() {
		fmt.Fprintf(&sb, "%v ", value)
	}

	sb.WriteString("]")

	return sb.String()
}
