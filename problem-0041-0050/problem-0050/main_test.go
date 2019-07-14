package main

import "testing"

func TestBinaryTreeEvaluation(t *testing.T) {
	tests := []struct {
		expr     *node
		expected int
	}{
		{n("*", n("+", l("3"), l("2")), n("+", l("4"), l("5"))), 45},
	}

	for i, test := range tests {
		actual := test.expr.evaluate()
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, test.expected, actual)
		}
	}
}

func n(val string, left, right *node) *node {
	res := newNode(val)
	res.left, res.right = left, right
	return res
}

func l(val string) *node {
	return newNode(val)
}
