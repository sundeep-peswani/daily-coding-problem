package main

import "testing"

func Test_AlternateLinkedList(t *testing.T) {
	tests := []struct{
		input []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 3, 2, 5, 4}},
	}

	for i, test := range tests {
		ll := newLinkedList(test.input)
		ll.alternate()
		actual := ll.toIntSlice()

		if !equal(test.expected, actual) {
			t.Errorf("Test %d: expected %v, actual %v\n", i + 1, test.expected, actual)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, c := range a {
		if b[i] != c {
			return false
		}
	}

	return true
}