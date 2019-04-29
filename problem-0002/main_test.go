package main

import "testing"

func TestArrayProduct(t *testing.T) {
	tables := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
		{[]int{3, 2, 1}, []int{2, 3, 6}},
	}

	for i, table := range tables {
		actual := arrayProduct(table.input)
		for j, n := range table.expected {
			if actual[j] != n {
				t.Errorf("Test %d: mismatch at index %d. Expected %d, actual %d", i, j, n, actual[j])
			}
		}
	}
}
