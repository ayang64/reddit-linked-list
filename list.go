package list

import "log"

type node struct {
	next  *node
	value interface{}
}

func Find(head *node, val interface{}, cmp func(interface{}, interface{}) bool) []interface{} {
	rc := []interface{}(nil)
	for v := head; v != nil; v = v.next {
		if cmp(val, v.value) {
			rc = append(rc, v.value)
		}
	}
	return rc
}

func Slice(head *node) []interface{} {
	rc := []interface{}(nil)

	for v := head; v != nil; v = v.next {
		rc = append(rc, v.value)
	}

	return rc
}

func Dump(head *node) {
	for v := head; v != nil; v = v.next {
		log.Printf("%#v", v)
	}
}

func Prepend(head **node, value interface{}) {
	*head = &node{next: *head, value: value}
}

func Append(head **node, value interface{}) {
	// find the end of the list
	cur := head
	for ; *cur != nil; cur = &(*cur).next {
	}

	// append to end of list
	*cur = &node{next: nil, value: value}
}
