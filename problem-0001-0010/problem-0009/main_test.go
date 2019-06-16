package main

import "testing"

func TestLargestSum(t *testing.T) {
	tables := []struct {
		arr      []int
		expected int
	}{
		{[]int{2, 4, 6, 2, 5}, 13},
		{[]int{2, 6, 4, 2, 5}, 11},
		{[]int{2, 6, 2, 4, 5}, 11},
		{[]int{6, 2, 4, 10, 5}, 16},
		{[]int{5, 1, 1, 5}, 10},
		{[]int{-1, -1, -1, -1}, -1},
		{[]int{-1, -2, -3, -4}, -1},
		{[]int{2, 4}, 4},
		{[]int{4, 2}, 4},
		{[]int{2}, 2},
		{[]int{}, 0},
	}

	for i, table := range tables {
		actual := largestSum(table.arr)
		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d", i, table.expected, actual)
		}
	}
}
