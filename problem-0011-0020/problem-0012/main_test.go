package main

import "testing"

func TestCountUniqueWaysToClimb(t *testing.T) {
	tables := []struct {
		n        int
		steps    []int
		expected int
	}{
		{0, []int{1, 2}, 0},
		{4, []int{1, 2}, 5},
		{3, []int{2, 4}, 0},
		{100, []int{1, 2}, 1298777728820984005},
		{100, []int{2, 5}, 545813094},
	}

	for i, table := range tables {
		memo := make(map[int]int)
		actual := countUniqueWaysToClimb(table.n, table.steps, memo)
		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i, table.expected, actual)
		}
	}
}
