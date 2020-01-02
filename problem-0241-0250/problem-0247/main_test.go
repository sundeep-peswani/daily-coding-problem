package main

import "testing"

func Test_IsHeightBalanced(t *testing.T) {
	tests := []struct{
		input *node
		expected bool
	}{
		{b(0, l(1, r(3, n(5))), r(2, b(4, n(6), n(7)))), true},
		{b(0, n(1), r(2, b(4, n(6), n(7)))), false},
	}

	for i, test := range tests {
		actual := test.input.isHeightBalanced()
		if actual != test.expected {
			t.Errorf("Test %d: expected %t, actual %t\n", i+1, test.expected, actual)
		}
	}
}

func b(v int, a, b *node) *node {
	return newNode(v, a, b)
}

func l(v int, a *node) *node {
	return newNode(v, a, nil)
}

func r(v int, b *node) *node {
	return newNode(v, nil, b)
}

func n(v int) *node {
	return newNode(v, nil, nil)
}