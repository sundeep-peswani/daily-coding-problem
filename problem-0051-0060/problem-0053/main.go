package main

import "fmt"

func main() {

}

type node struct {
	value int
	next  *node
}

func newNode(val int) *node {
	return &node{val, nil}
}

type stack struct {
	head *node
}

func (s *stack) isEmpty() bool {
	return s.head == nil
}

func (s *stack) peek() (int, error) {
	if s.isEmpty() {
		return -1, fmt.Errorf("stack is empty")
	}

	return s.head.value, nil
}

func (s *stack) pop() (int, error) {
	if s.isEmpty() {
		return -1, fmt.Errorf("stack is empty")
	}

	res := s.head.value
	s.head = s.head.next
	return res, nil
}

func (s *stack) push(n int) {
	next := newNode(n)
	if s.isEmpty() {
		s.head = next
	} else {
		next.next = s.head
		s.head = next
	}
}

func (s *stack) migrateTo(o *stack) {
	for !s.isEmpty() {
		if n, err := s.pop(); err == nil {
			o.push(n)
		}
	}
}

type queue struct {
	forward, backward *stack
}

func newQueue() *queue {
	return &queue{&stack{nil}, &stack{nil}}
}

func (q *queue) enqueue(n int) {
	if q.backward.isEmpty() {
		q.forward.migrateTo(q.backward)
	}

	q.backward.push(n)
}

func (q *queue) dequeue() (int, error) {
	if n, err := q.peek(); err != nil {
		return -1, err
	} else {
		q.forward.pop()
		return n, nil
	}
}

func (q *queue) peek() (int, error) {
	if q.forward.isEmpty() {
		q.backward.migrateTo(q.forward)
	}

	if n, err := q.forward.peek(); err != nil {
		return -1, fmt.Errorf("queue is empty")
	} else {
		return n, nil
	}
}
