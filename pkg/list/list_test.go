package list_test

import (
	"testing"

	"github.com/1995parham/linkedlist/pkg/list"
)

func TestListWith3Member(t *testing.T) {
	l := list.New[int]()
	l.PushFront(10)

	if l.Len() != 1 {
		t.Errorf("list should have %d items in it", l.Len())
	}
}
