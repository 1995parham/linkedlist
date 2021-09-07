package list

import "testing"

func TestListWith3Member(t *testing.T) {
	l := New[int]()
	l.PushFront(10)
	if len(l) != 1 {
		t.Errorf("list should have %d items in it", len(l))	
	}
}
