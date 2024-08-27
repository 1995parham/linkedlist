package list

import (
	"testing"
)

func TestListWith2Members(t *testing.T) {
	l := New[int]()

	l.PushFront(10)
	l.PushFront(20)

	if l.Len() != 2 {
		t.Errorf("list should have 2 items but have %d items in it", l.Len())
	}

	first := l.Head
	if first.data() != 20 {
		t.Errorf("first item should be 20")
	}

	if first.next().data() != 10 {
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

	first := l.Head
	if first.data() != 10 {
		t.Errorf("first item should be 10")
	}

	if first.next().data() != 20 {
		t.Errorf("next item should be 20")
	}
}

func TestFilter(t *testing.T) {
	l := New[int]()

	l.PushBack(10)
	l.PushBack(11)
	l.PushBack(20)
	l.PushBack(21)

	r := l.Filter(func(i int) bool {
		return i%2 == 0
	})

	resultLen := 0
	resultSet := make(map[int]bool)

	for i := range r {
		resultSet[i] = true
		resultLen += 1
	}

	if resultLen != 2 {
		t.Errorf("result after filtering must contains only 2 items")
	}

	if !resultSet[10] {
		t.Errorf("result after filtering must contains 10")
	}

	if !resultSet[20] {
		t.Errorf("result after filtering must contains 20")
	}
}

func TestString(t *testing.T) {
	l := New[int]()

	l.PushBack(10)
	l.PushBack(11)
	l.PushBack(20)
	l.PushBack(21)

	if l.String() != "[ 10 11 20 21 ]" {
		t.Errorf("linkedlist string representation is not currect")
	}
}

func TestListIsLinker(t *testing.T) {
	t.Parallel()

	// test for interface implementation
	var _ Linker[interface{}] = &List[interface{}]{}

	var _ Linker[interface{}] = New[interface{}]()
}
