package main

import (
	"fmt"
	"slices"

	"github.com/1995parham/linkedlist/pkg/list"
)

// Int demonstrates implementing Comparable[C Comparable[C]] (Go 1.26 recursive generics).
type Int int

func (a Int) Less(b Int) bool {
	return a < b
}

func main() {
	l := list.New[int]()

	l.PushFront(10)
	l.PushFront(20)
	l.PushFront(40)
	l.PushFront(42)
	l.PushBack(1378)

	fmt.Println(l)

	// iter.Seq[T] Values iterator (Go 1.23 convention)
	for value := range l.Values() {
		fmt.Printf("- %d\n", value)
	}

	// iter.Seq2[int, T] All iterator with index-value pairs (Go 1.23 convention)
	for i, value := range l.All() {
		fmt.Printf("[%d] %d\n", i, value)
	}

	// Collect materializes the list into a slice using slices.Collect (Go 1.23)
	items := l.Collect()
	fmt.Println("collected:", items)

	// Filter returns an iter.Seq[T], collect it into a slice with slices.Collect
	evens := slices.Collect(l.Filter(func(i int) bool {
		return i%10 == 0
	}))
	fmt.Println("evens:", evens)

	// Map is a standalone generic function with two type parameters [T, U any]
	// demonstrating multi-type-parameter generics (methods can't have extra type params)
	doubled := slices.Collect(list.Map(l.Values(), func(i int) int {
		return i * 2
	}))
	fmt.Println("doubled:", doubled)

	// Map can also change the element type (int -> string)
	labels := slices.Collect(list.Map(l.Values(), func(i int) string {
		return fmt.Sprintf("#%d", i)
	}))
	fmt.Println("labels:", labels)

	// SortedList uses Go 1.26 recursive generic types:
	//   type Comparable[C Comparable[C]] interface { Less(C) bool }
	// Elements are automatically inserted in ascending order.
	sl := list.NewSorted[Int]()

	sl.Insert(30)
	sl.Insert(10)
	sl.Insert(50)
	sl.Insert(20)
	sl.Insert(40)

	fmt.Println("sorted:", sl)
	fmt.Println("sorted collected:", sl.Collect())
}
