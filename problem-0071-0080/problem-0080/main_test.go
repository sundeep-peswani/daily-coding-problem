package main

import "testing"

func TestFindDeepestNode(t *testing.T) {
	tests := []struct {
		tree     *node
		expected int
	}{
		{b(1, l(2, n(4)), n(3)), 4},
		{b(1, l(2, l(4, l(5, l(6, l(7, n(8)))))), n(3)), 8},
		{b(1, l(2, n(4)), r(3, n(5))), 5},
		{b(1, l(2, n(4)), r(3, r(5, n(6)))), 6},
		{b(1, l(2, n(4)), r(3, l(5, n(6)))), 6},
	}

	for i, test := range tests {
		actual := findDeepestNode(test.tree)
		if actual == nil {
			t.Errorf("Test %d: expected %d, actual nil\n", i+1, test.expected)
		}
		if actual.value != test.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, test.expected, actual.value)
		}
	}
}

func b(val int, left, right *node) *node {
	return &node{val, left, right}
}

func l(val int, left *node) *node {
	return &node{val, left, nil}
}

func r(val int, right *node) *node {
	return &node{val, nil, right}
}

func n(val int) *node {
	return &node{val, nil, nil}
}
