package list

import (
	"slices"
	"testing"
)

type Int int

func (a Int) Less(b Int) bool {
	return a < b
}

func TestSortedInsertUnordered(t *testing.T) {
	sl := NewSorted[Int]()

	sl.Insert(30)
	sl.Insert(10)
	sl.Insert(20)

	result := sl.Collect()
	expected := []Int{10, 20, 30}

	if !slices.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestSortedInsertDuplicates(t *testing.T) {
	sl := NewSorted[Int]()

	sl.Insert(10)
	sl.Insert(10)
	sl.Insert(5)
	sl.Insert(10)

	result := sl.Collect()
	expected := []Int{5, 10, 10, 10}

	if !slices.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestSortedInsertDescending(t *testing.T) {
	sl := NewSorted[Int]()

	sl.Insert(50)
	sl.Insert(40)
	sl.Insert(30)
	sl.Insert(20)
	sl.Insert(10)

	result := sl.Collect()
	expected := []Int{10, 20, 30, 40, 50}

	if !slices.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestSortedInsertAscending(t *testing.T) {
	sl := NewSorted[Int]()

	sl.Insert(10)
	sl.Insert(20)
	sl.Insert(30)

	result := sl.Collect()
	expected := []Int{10, 20, 30}

	if !slices.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestSortedInsertSingle(t *testing.T) {
	sl := NewSorted[Int]()

	sl.Insert(42)

	if sl.Len() != 1 {
		t.Errorf("expected length 1, got %d", sl.Len())
	}

	result := sl.Collect()
	if !slices.Equal(result, []Int{42}) {
		t.Errorf("got %v, want [42]", result)
	}
}

func TestSortedString(t *testing.T) {
	sl := NewSorted[Int]()

	sl.Insert(30)
	sl.Insert(10)
	sl.Insert(20)

	if sl.String() != "[ 10 20 30 ]" {
		t.Errorf("got %q, want %q", sl.String(), "[ 10 20 30 ]")
	}
}

func TestSortedFilter(t *testing.T) {
	sl := NewSorted[Int]()

	sl.Insert(15)
	sl.Insert(10)
	sl.Insert(25)
	sl.Insert(20)

	evens := slices.Collect(sl.Filter(func(i Int) bool {
		return i%10 == 0
	}))

	expected := []Int{10, 20}

	if !slices.Equal(evens, expected) {
		t.Errorf("got %v, want %v", evens, expected)
	}
}
