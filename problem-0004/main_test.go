package main

import "testing"

func TestFindFirstMissingPositiveInteger(t *testing.T) {
	tables := []struct {
		input    []int
		expected int
	}{
		{[]int{3, 4, -1, 1}, 2},
		{[]int{1, 2, 0}, 3},
		{[]int{2, 1, 0}, 3},
		{[]int{1, 2, 3, 0}, 4},
		{[]int{0, 1, 2, 3, 5}, 4},
		{[]int{-4, -3, -2, -1}, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 2},
		{[]int{1, 2, 3, 4, 5}, 6},
	}

	for i, table := range tables {
		actual := findFirstMissingPositiveInteger(table.input)

		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d", i, table.expected, actual)
		}
	}
}
