package main

import "fmt"

func main() {

}

func usage() {

}

func findIntersection(a, b *linkedList) *node {
	seen := make(map[string]bool)
	n := a.head

	for n != nil {
		seen[n.addr()] = true
		n = n.next
	}

	n = b.head
	for n != nil {
		if _, ok := seen[n.addr()]; ok {
			return n
		}
		n = n.next
	}

	return nil
}

type node struct {
	val  int
	next *node
}

func newNode(v int) *node {
	return &node{v, nil}
}

func (n *node) String() string {
	return fmt.Sprintf("node(%p): %d -> node(%p)", n, n.val, n.next)
}

func (n *node) addr() string {
	return fmt.Sprintf("%p", n)
}

type linkedList struct {
	head *node
	tail *node
}

func newLinkedList() *linkedList {
	return &linkedList{nil, nil}
}

func (l *linkedList) String() string {
	return fmt.Sprintf("list(%p): head = %p", l, l.head)
}

func (l *linkedList) push(n *node) {
	if l.head == nil && l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
}
