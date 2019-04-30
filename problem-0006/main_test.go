package main

import "testing"

func TestNewNode(t *testing.T) {
	tables := []struct {
		value string
		prev  int64
	}{
		{"first", 0},
		{"second", 1},
		{"third", 2},
	}

	h := newHeap()

	for i, table := range tables {
		if _, err := h.newNode(table.value, table.prev); err != nil {
			t.Errorf("Test %d: unexpected error: %v\n", i, err)
		}
	}
}

func TestAddAndGet(t *testing.T) {
	values := []string{"first", "second", "third", "fourth"}

	h := newHeap()
	head, err := h.newNode("head", 0)
	if err != nil {
		t.Error(err)
	}

	for i, value := range values {
		if err := head.add(value); err != nil {
			t.Errorf("Test %d: unable to add %s: %v\n", i, value, err)
		}
	}

	for i, value := range values {
		if actual, err := head.get(i + 1); err != nil {
			t.Errorf("Test %d: unable to get item at %d: %v\n", i, i+1, err)
		} else if actual.value != value {
			t.Errorf("Test %d: expected %s, actual %s\n", i, value, actual.value)
		}
	}
}
