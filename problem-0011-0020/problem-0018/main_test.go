package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMaxSubarray(t *testing.T) {
	tables := []struct {
		array    []int
		k        int
		expected []int
	}{
		{[]int{10, 5, 2, 7, 8, 7}, 3, []int{10, 7, 8, 8}},
		{[]int{10, 5, 2, 7, 8, 7}, 2, []int{10, 5, 7, 8, 8}},
		{[]int{10, 5, 2, 7, 8, 7}, 1, []int{10, 5, 2, 7, 8, 7}},
		{[]int{10, 5, 2, 7, 8, 7}, 6, []int{10}},
		{[]int{10, 5, 2, 7, 8, 7}, 0, []int{}},
		{[]int{}, 6, []int{}},
	}

	for i, table := range tables {
		actual := maxSubarray(table.array, table.k)
		if !reflect.DeepEqual(table.expected, actual) {
			t.Errorf("Test %d: expected %#v, actual %#v\n", i+1, table.expected, actual)
		}
	}
}

func TestDeque(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	d := &deque{nil, nil}

	if !d.isEmpty() {
		t.Errorf("Expected queue to be empty, actually isn't")
	}

	n := 2 + rand.Intn(10)

	for i := 0; i < n; i++ {
		r := rand.Intn(20)

		if i%2 == 0 {
			d.push(&node{-1 - r, nil, nil})
		} else {
			d.pushBack(&node{1 + r, nil, nil})
		}
	}

	if d.isEmpty() {
		t.Errorf("Filled queue with %d items, is empty", n)
	}

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			if d.peek() == nil {
				t.Errorf("Queue is prematurely empty, expected %d more items\n", n-i)
			} else if d.peek().val > 0 {
				t.Errorf("Queue was not filled correctly, expected <0, actual %d\n", d.peek().val)
			}
		} else {
			if d.peekBack() == nil {
				t.Errorf("Queue is prematurely empty, expected %d more items\n", n-i)
			} else if d.peekBack().val < 0 {
				t.Errorf("Queue was not filled correctly, expected >0, actual %d\n", d.peekBack().val)
			}
		}

		if d.isEmpty() {
			t.Errorf("Queue should still have %d items, is empty", n-i)
		}
	}

	var v *node
	for i := 0; i < n; i++ {
		v = nil

		if i%2 == 0 {
			v = d.pop()
		} else {
			v = d.popBack()
		}

		if v == nil {
			t.Errorf("Popping %d: failed to pop, received nil\n", i)
		}
	}

	if !d.isEmpty() {
		t.Errorf("Queue should be empty, is not: %s\n", d)
	}
}
