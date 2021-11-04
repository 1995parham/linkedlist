package list

import (
	"testing"
	)

func TestListWith2Members(t *testing.T) {
	var l *List[int] = New[int]()
	l.PushFront(10)
	l.PushFront(20)

	if l.Len() != 2 {
		t.Errorf("list should have 2 items but have %d items in it", l.Len())
	}
	first := l.Head.(Node[int])
	if getData[int](first) != 20 {
		t.Errorf("first item should be 20")
	}
	if getData[int](first.next()) != 10 {
		t.Errorf("next item should be 10")
	}
}


func TestListWith2MembersBack(t *testing.T) {
	l := New[int]()
	l.PushBack(10)
	l.PushBack(20)

	if l.Len() != 2 {
		t.Errorf("list should have 2 items but have %d items in it", l.Len())
	}
	first := l.Head.(Node[int])
	if getData[int](first) != 10 {
		t.Errorf("first item should be 10")
	}
	if getData[int](first.next()) != 20 {
		t.Errorf("next item should be 20")
	}
}


func TestListIsLinker(t *testing.T) {
	//test for interface implementation
	var _ Linker[interface{}] = &List[interface{}]{}
	var _ Linker[interface{}] = New[interface{}]()
}