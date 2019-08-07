package list

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	var head *node = nil

	Prepend(&head, "test1")
	Prepend(&head, "test2")
	Prepend(&head, "test3")
	Append(&head, "needle2")
	Prepend(&head, "test4")
	Prepend(&head, "test5")
	Append(&head, "test1")
	Append(&head, "needle1")
	Append(&head, "test2")
	Append(&head, "test3")
	Append(&head, "test4")
	Append(&head, "needle3")
	Append(&head, "test5")

	cmp := func(a interface{}, b interface{}) bool {
		// we only compare strings

		t.Logf("here a")
		s1, ok := a.(string)
		if !ok {
			return false
		}

		t.Logf("here b")
		s2, ok := b.(string)
		if !ok {
			t.Logf("here c")
			return false
		}
		t.Logf("here c")

		t.Logf("comparing %q with %q", s1, s2)
		if strings.HasPrefix(s1, s2) {
			return true
		}

		return false
	}

	vals := Find(head, "needle", cmp)

	t.Logf("vals = %v", vals)

}

func TestSlice(t *testing.T) {
	var head *node = nil

	Prepend(&head, "test1")
	Prepend(&head, "test2")
	Prepend(&head, "test3")
	Prepend(&head, "test4")
	Prepend(&head, "test5")
	Append(&head, "test1")
	Append(&head, "test2")
	Append(&head, "test3")
	Append(&head, "test4")
	Append(&head, "test5")

	Dump(head)

	s := Slice(head)

	t.Logf("s = %v", s)
}

func TestPrepend(t *testing.T) {
	var head *node = nil

	Prepend(&head, "test1")
	Prepend(&head, "test2")
	Prepend(&head, "test3")
	Prepend(&head, "test4")
	Prepend(&head, "test5")

	Dump(head)
}

func TestAppend(t *testing.T) {
	var head *node = nil

	Append(&head, "test1")
	Append(&head, "test2")
	Append(&head, "test3")
	Append(&head, "test4")
	Append(&head, "test5")

	Dump(head)
}
