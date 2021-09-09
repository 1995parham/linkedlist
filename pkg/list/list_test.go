package list

import "testing"

func TestListWith3Member(t *testing.T) {
	l := New[int]()
	l.PushFront(10)

	if l.Len() != 1 {
		t.Errorf("list should have %d items in it", l.Len())
	}
}
