package main

import "testing"

func TestCalculateWaterStorage(t *testing.T) {
	tests := []struct {
		elevationMap []int
		expected     int
	}{
		{[]int{2, 1, 2}, 1},
		{[]int{3, 0, 1, 3, 0, 5}, 8},
		{[]int{3, 3, 2, 0}, 0},
		{[]int{0, 1, 1, 3}, 0},
		{[]int{3, 0, 1, 2, 3}, 6},
		{[]int{3, 5, 0, 4}, 4},
		{[]int{3}, 0},
	}

	for i, test := range tests {
		actual := calcuateWaterStorage(test.elevationMap)
		if actual != test.expected {
			t.Errorf("Test %d: expacted %d, actual %d: %v\n", i+1, test.expected, actual, test.elevationMap)
		}
	}
}
