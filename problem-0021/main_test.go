package main

import "testing"

func TestCountRooms(t *testing.T) {
	tables := []struct {
		times    []int
		expected int
	}{
		{[]int{30, 75, 0, 50, 60, 150}, 2},
		{[]int{0, 1, 2, 3, 4, 5}, 1},
		{[]int{0, 3, 0, 1, 0, 2}, 3},
		{[]int{0, 1, 1, 2, 2, 3}, 1},
		{[]int{0, 1, 1, 2, 0, 2}, 2},
		{[]int{3, 5, 4, 6, 1, 4}, 2},
		{[]int{1, 4, 4, 6, 5, 7, 4, 5}, 2},
	}

	for i, table := range tables {
		actual := countRooms(createSchedule(table.times))

		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i, table.expected, actual)
		}
	}
}
