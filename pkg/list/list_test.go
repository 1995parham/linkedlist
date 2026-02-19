package list

import (
	"fmt"
	"slices"
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

	result := slices.Collect(l.Filter(func(i int) bool {
		return i%2 == 0
	}))

	if len(result) != 2 {
		t.Errorf("result after filtering must contains only 2 items")
	}

	if !slices.Contains(result, 10) {
		t.Errorf("result after filtering must contains 10")
	}

	if !slices.Contains(result, 20) {
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

func TestCollect(t *testing.T) {
	l := New[int]()

	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	result := l.Collect()

	if !slices.Equal(result, []int{10, 20, 30}) {
		t.Errorf("Collect() = %v, want [10 20 30]", result)
	}
}

func TestValues(t *testing.T) {
	l := New[int]()

	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	result := slices.Collect(l.Values())

	if !slices.Equal(result, []int{10, 20, 30}) {
		t.Errorf("Values() = %v, want [10 20 30]", result)
	}
}

func TestAll(t *testing.T) {
	l := New[int]()

	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	indices := make([]int, 0)
	values := make([]int, 0)

	for i, v := range l.All() {
		indices = append(indices, i)
		values = append(values, v)
	}

	if !slices.Equal(indices, []int{0, 1, 2}) {
		t.Errorf("All() indices = %v, want [0 1 2]", indices)
	}

	if !slices.Equal(values, []int{10, 20, 30}) {
		t.Errorf("All() values = %v, want [10 20 30]", values)
	}
}

func TestMap(t *testing.T) {
	l := New[int]()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	doubled := slices.Collect(Map(l.Values(), func(i int) int {
		return i * 2
	}))

	if !slices.Equal(doubled, []int{2, 4, 6}) {
		t.Errorf("Map() = %v, want [2 4 6]", doubled)
	}
}

func TestMapTypeChange(t *testing.T) {
	l := New[int]()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	strings := slices.Collect(Map(l.Values(), func(i int) string {
		return fmt.Sprintf("#%d", i)
	}))

	if !slices.Equal(strings, []string{"#1", "#2", "#3"}) {
		t.Errorf("Map() type change = %v, want [#1 #2 #3]", strings)
	}
}

func TestListIsLinker(t *testing.T) {
	t.Parallel()

	// test for interface implementation using any (alias for interface{})
	var _ Linker[any] = &List[any]{}

	var _ Linker[any] = New[any]()
}
