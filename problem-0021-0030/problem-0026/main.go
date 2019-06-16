package main

func main() {

}

type linkedList struct {
	head, tail *node
}

type node struct {
	value int
	next  *node
}

func newLinkedList() *linkedList {
	return &linkedList{nil, nil}
}

func (l *linkedList) isEmpty() bool {
	return l.head == nil && l.tail == nil
}

func (l *linkedList) pushBack(value int) {
	n := newNode(value)

	if l.isEmpty() {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
}

func (l *linkedList) remove(k int) bool {
	start := l.head

	current := l.head
	for i := 0; i <= k; i++ {
		if current == nil {
			return false
		}
		current = current.next
	}

	for current != nil {
		current = current.next
		start = start.next
	}

	start.next = start.next.next

	return true
}

func newNode(value int) *node {
	return &node{value, nil}
}
