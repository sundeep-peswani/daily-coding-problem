package main

import "testing"

func TestAdvanceToEnd(t *testing.T) {
	tests := []struct {
		input []int
		expected bool
	}{
		{[]int{1, 3, 1, 2, 0, 1}, true},
		{[]int{1, 2, 1, 0, 0}, false},
		{[]int{1}, true},
		{[]int{57}, true},
		{[]int{400, 200}, true},
		{[]int{0, 0, 0, 0, 0, 300}, false},
		{[]int{0, 0, 0, 0, 3}, false},
		{[]int{1, 67, 0, 0, 4}, true},
	}

	for i, test := range tests {
		actual := advanceToEnd(test.input)

		if actual != test.expected {
			t.Errorf("Test %d: expected %v, actual %v: %v\n", i+1, test.expected, actual, test.input)
		}
	}
}