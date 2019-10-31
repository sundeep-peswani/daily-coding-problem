package main

import "testing"

func TestFindSecondLargest(t *testing.T) {
	tests := []struct{
		input []int
		expected int
	}{
		{[]int{1,2}, 1},
		{[]int{1,2,3}, 2},
		{[]int{1,2,3,4}, 3},
		{[]int{1,2,3,4,5}, 4},
		{[]int{1,2,3,4,5,6}, 5},
		{[]int{1,2,3,4,5,6,7}, 6},
	}

	for i, test := range tests {
		actual := findSecondLargest(createBST(test.input))
		if test.expected != actual {
			t.Errorf("Test %d: expected %d, actual %d for: %v\n", i+1, test.expected, actual, test.input)
		}
	}
}