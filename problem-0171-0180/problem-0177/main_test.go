package main

import (
	"testing"
)

func TestRotateLinkedList(t *testing.T) {
	tests := []struct {
		start *list
		k int
		expected *list
	}{
		{generateList([]int{7, 7, 3, 5}), 2, generateList([]int{3, 5, 7, 7})},
		{generateList([]int{1, 2, 3, 4, 5}), 3, generateList([]int{4, 5, 1, 2, 3})}, // README is wrong
		{generateList([]int{7, 7, 3, 5}), 6, generateList([]int{3, 5, 7, 7})},
		{generateList([]int{1, 2, 3, 4, 5}), 0, generateList([]int{1, 2, 3, 4, 5})},
	}

	for i, test := range tests {
		test.start.rotate(test.k)

		if (!test.start.equals(test.expected)) {
			t.Errorf("Test %d: expected %s, got %s\n", i+1, test.expected, test.start)
		}
	}
}

func generateList(arr []int) *list {
	l := &list{nil, nil}

	for i := 0; i < len(arr); i++ {
		l.append(arr[i])
	}

	return l
}
