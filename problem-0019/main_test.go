package main

import "testing"

func TestCalculateMinimumCost(t *testing.T) {
	tables := []struct {
		costMatrix [][]int
		expected   int
	}{
		{[][]int{[]int{1, 2, 3, 4, 5, 6}}, 1},
		{[][]int{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}}, 2},
		{[][]int{[]int{1, 2, 3}, []int{3, 2, 1}, []int{4, 9, 4}}, 6},
	}

	for i, table := range tables {
		actual := calculateMinimumCost(table.costMatrix)
		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i, table.expected, actual)
		}
	}
}
