package main

func main() {

}

type linkedList struct {
	head, tail *node
}

func newLinkedList(input []int) *linkedList {
	var ll linkedList

	for _, item := range input {
		ll.push(newNode(item))
	}

	return &ll
}

func (l *linkedList) push(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	l.tail = n
	return
}

func (l *linkedList) alternate() {
	if l.head == nil {
		return
	}

	a, b := l.head, l.head.next
	for a != nil {
		if b == nil {
			return
		}

		c := b.next
		a.next = c
		b.next = c.next
		c.next = b

		a, b = b, b.next
	}
}

func (l linkedList) toIntSlice() []int {
	var result []int

	curr := l.head
	for curr != nil {
		result = append(result, curr.value)
		curr = curr.next
	}

	return result
}

type node struct {
	value int
	next *node
}

func newNode(val int) *node {
	return &node{val, nil}
}