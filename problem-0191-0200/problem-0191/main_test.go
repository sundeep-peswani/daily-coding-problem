package main

import "testing"

func TestInterleavingIntervals(t *testing.T) {
	tests := []struct{
		input []interval
		expected int
	}{
		{[]interval{create(7, 9), create(2, 4), create(5, 8)}, 1},
		{[]interval{create(0, 1), create(1, 2)}, 0},
	}

	for i, test := range tests {
		actual := interleavingIntervals(test.input)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d: %v\n", i+1, test.expected, actual, test.input)
		}
	}
}

func create(start, end int) interval {
	return interval{start, end}
}