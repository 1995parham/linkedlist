package main

import (
	"fmt"

	"github.com/1995parham/linkedlist/pkg/list"
)

func main() {
	l := list.New[int]()

	l.PushFront(10)
	l.PushFront(20)
	l.PushFront(40)
	l.PushFront(42)
	l.PushBack(1378)

	fmt.Println(l)

	s := l.Filter(func(i int) bool {
		return i%10 == 0
	})

	fmt.Println(s)
}
