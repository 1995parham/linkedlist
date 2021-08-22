package list

import "testing"

func TestListWith3Member(t *testing.T) {
	l := New[int]()

	l.PushFront(10)
}
