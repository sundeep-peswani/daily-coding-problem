package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestStack(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	s := &stack{nil}

	if v, err := s.peek(); err == nil {
		t.Fatalf("stack.peek(): expected error on empty stack, got %d\n", v)
	}
	if v, err := s.pop(); err == nil {
		t.Fatalf("stack.pop(): expected error on empty stack, got %d\n", v)
	}

	arr := generateArray(10000)
	for _, n := range arr {
		s.push(n)
		if v, err := s.peek(); err != nil {
			t.Fatalf("stack.push(): unable to push to stack, got error: %v\n", err)
		} else if v != n {
			t.Fatalf("stack.push(): expected %d, actual %d\n", n, v)
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if n, err := s.peek(); err != nil {
			t.Errorf("stack.peek(%d): expected %d, got error: %v\n", i, arr[i], err)
		} else if n != arr[i] {
			t.Errorf("stack.peek(%d): expected %d, got %d\n", i, arr[i], n)
		}

		if n, err := s.pop(); err != nil {
			t.Errorf("stack.pop(%d): expected %d, got error: %v\n", i, arr[i], err)
		} else if n != arr[i] {
			t.Errorf("stack.pop(%d): expected %d, got %d\n", i, arr[i], n)
		}
	}

	if v, err := s.peek(); err == nil {
		t.Fatalf("stack.peek(): expected error on empty stack, got %d\n", v)
	}
	if v, err := s.pop(); err == nil {
		t.Fatalf("stack.pop(): expected error on empty stack, got %d\n", v)
	}
}

func TestQueue(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	q := newQueue()

	if v, err := q.peek(); err == nil {
		t.Fatalf("queue.peek(): expected error on empty queue, got %d\n", v)
	}
	if v, err := q.dequeue(); err == nil {
		t.Fatalf("queue.dequeue(): expected error on empty queue, got %d\n", v)
	}

	arr := generateArray(10000)
	for _, n := range arr {
		q.enqueue(n)
	}

	for i, n := range arr {
		if v, err := q.peek(); err != nil {
			t.Errorf("queue.peek(%d): expected %d, got error: %v\n", i, arr[i], err)
		} else if n != arr[i] {
			t.Errorf("queue.peek(%d): expected %d, got %d\n", i, n, v)
		}

		if v, err := q.dequeue(); err != nil {
			t.Errorf("queue.dequeue(%d): expected %d, got error: %v\n", i, arr[i], err)
		} else if v != n {
			t.Errorf("queue.dequeue(%d): expected %d, got %d\n", i, n, v)
		}
	}

	if v, err := q.peek(); err == nil {
		t.Fatalf("queue.peek(): expected error on empty queue, got %d\n", v)
	}
	if v, err := q.dequeue(); err == nil {
		t.Fatalf("queue.dequeue(): expected error on empty queue, got %d\n", v)
	}
}

func generateArray(len int) []int {
	var res = make([]int, len)

	for i := 0; i < len; i++ {
		res[i] = rand.Intn(len)
	}

	return res
}
