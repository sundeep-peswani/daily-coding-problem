package main

import (
	"testing"
)

func TestCountUnival(t *testing.T) {
	tables := []struct {
		root     *Node
		expected int
	}{
		{nil, 0},
		{n(0), 1},
		{nl(0, n(1)), 1},
		{nr(0, n(1)), 1},
		{nl(0, nr(1, n(1))), 1},
		{nb(1, n(1), n(1)), 3},
		{nb(1, nb(1, n(1), n(1)), nb(1, n(1), n(1))), 7},
		{nb(1, nb(1, n(0), n(1)), nb(1, n(1), n(0))), 4},
		{nb(1, nb(1, n(0), n(1)), nb(1, n(1), n(1))), 5},
	}

	for i, table := range tables {
		actual, _ := table.root.countUnival()
		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i, table.expected, actual)
		}
	}
}

func nb(val int, left, right *Node) *Node {
	return &Node{val, left, right}
}

func nl(val int, left *Node) *Node {
	return nb(val, left, nil)
}

func nr(val int, right *Node) *Node {
	return nb(val, nil, right)
}

func n(val int) *Node {
	return nb(val, nil, nil)
}
