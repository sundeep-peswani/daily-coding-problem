package main

import "testing"

func TestFindSum(t *testing.T) {
	tree := &binaryNode{10, &binaryNode{5, nil, nil}, &binaryNode{15, &binaryNode{11, nil, nil}, &binaryNode{15, nil, nil}}}

	tests := []struct{
		k int
		expected bool
	}{
		{20, true},
		{26, true},
		{15, true},
		{3, false},
		{40, false},
		{16, true},
		{21, true},
		{22, false},
	}

	for i, test := range tests {
		actual := findSum(tree, test.k)
		if test.expected && actual == nil {
			t.Errorf("Test %d: expected a result for %d, got nothing\n", i + 1, test.k)
		} else if !test.expected && actual != nil {
			t.Errorf("Test %d: did not expect a result for %d, actual: %d, %d\n", i + 1, test.k, actual[0].value, actual[1].value)
		}
	}	
}