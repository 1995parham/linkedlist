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

	fmt.Println(l)
}
