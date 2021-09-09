package list_test

import (
	"testing"

	"github.com/1995parham/linkedlist/pkg/list"
)

func TestListWith3Member(t *testing.T) {
	l := list.New[int]()
	l.PushFront(10)
	l.PushFront(20)

	if l.Len() != 2 {
		t.Errorf("list should have 2 items but have %d items in it", l.Len())
	}
}
