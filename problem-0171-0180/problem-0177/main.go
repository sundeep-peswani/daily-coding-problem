package main

import (
	"fmt"
	"strings"
)

func main() {

}

type node struct {
	val int
	next *node
}

type list struct {
	head, tail *node
}

func (l *list) append(val int) {
	n := &node{val, nil}

	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	l.tail = n
}

func (l *list) len() int {
	n := l.head
	c := 0
	for n != nil {
		c++
		n = n.next
	}

	return c
}

func (l *list) rotate(k int) {
	steps := k % l.len()

	if steps == 0 {
		return
	}

	oldHead := l.head
	prevHead := l.head
	newHead := l.head

	for i := 0; i < steps; i++ {
		prevHead = newHead
		newHead = newHead.next
	}

	l.head = newHead
	l.tail.next = oldHead
	prevHead.next = nil
}

func (l *list) equals(m *list) bool {
	ln, mn := l.head, m.head

	for ln != nil && mn != nil && ln.val == mn.val {
		ln = ln.next
		mn = mn.next
	}

	return ln == nil && mn == nil
}

func (l *list) String() string {
	var b strings.Builder

	n := l.head

	for n != nil {
		b.WriteString(fmt.Sprintf("%d ", n.val))
		if n.next != nil {
			b.WriteString("> ")
		}
		n = n.next
	}

	return b.String()
}
