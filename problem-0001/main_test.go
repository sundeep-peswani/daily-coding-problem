package main

import "testing"

func TestFindSum(t *testing.T) {
	tables := []struct {
		total    int
		arr      []int
		expected bool
	}{
		{17, []int{10, 15, 3, 7}, true},
		{16, []int{10, 15, 3, 7}, false},
		{0, []int{}, false},
		{0, []int{-1, 1, -2, 2, 0, 0}, true},
		{2, []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}, true},
		{2, []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 0}, true},
	}

	for i, table := range tables {
		actual := findSum(table.total, table.arr)
		if actual != table.expected {
			t.Errorf("Test %d: expected %t, actual %t", i, table.expected, actual)
		}
	}
}
