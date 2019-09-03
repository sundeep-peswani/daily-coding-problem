package main

import "testing"

func TestNextBiggest(t *testing.T) {
	tests := []struct{
		bst *node
		search int
		expected int
	}{
		{b(10, n(5), b(30, n(22), n(35))), 22, 30},
		{b(10, n(5), b(30, n(22), n(35))), 10, 22},
		{b(10, n(5), b(30, n(22), n(35))), 30, 35},
		{b(5, b(2, n(-4), n(3)), b(21, n(19), r(25, n(26)))), -4, 2},
		{b(5, b(2, n(-4), n(3)), b(21, n(19), r(25, n(26)))), 2, 3},
		{b(5, b(2, n(-4), n(3)), b(21, n(19), r(25, n(26)))), 3, 5},
		{b(5, b(2, n(-4), n(3)), b(21, n(19), r(25, n(26)))), 5, 19},
		{b(5, b(2, n(-4), n(3)), b(21, n(19), r(25, n(26)))), 19, 21},
		{b(5, b(2, n(-4), n(3)), b(21, n(19), r(25, n(26)))), 21, 25},
		{b(5, b(2, n(-4), n(3)), b(21, n(19), r(25, n(26)))), 25, 26},
	}

	for i, test := range tests {
		actual := nextBiggest(test.bst, test.search)
		if actual != test.expected {
			t.Errorf("Test %d: searched for next biggest after %d, expected %d, actual %d\n", i+1, test.search, test.expected, actual)
		}
	}
}

func n(value int) *node {
	return &node{value, nil, nil, nil}
}

func l(value int, left *node) *node {
	newNode := &node{value, nil, left, nil}
	left.parent = newNode
	return newNode
}

func r(value int, right *node) *node {
	newNode := &node{value, nil, nil, right}
	right.parent = newNode
	return newNode
}

func b(value int, left, right *node) *node {
	newNode := &node{value, nil, left, right}
	left.parent, right.parent = newNode, newNode
	return newNode
}