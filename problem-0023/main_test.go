package main

import (
	"fmt"
	"testing"
)

func TestSolveMaze(t *testing.T) {
	tables := []struct {
		maze       string
		start, end []int
		expected   int
		err        error
	}{
		{"....,##.#,....,....", []int{3, 0}, []int{0, 0}, 7, nil},
		{".##.,.#..,.#..,...#", []int{0, 0}, []int{0, 3}, 9, nil},
		{".##.,####,....,....", []int{0, 0}, []int{0, 3}, 0, fmt.Errorf("Unable to solve")},
	}

	for i, table := range tables {
		actual, err := newMaze(table.maze).solve(table.start, table.end)

		if err != nil && table.err == nil {
			t.Errorf("Test %d: expected %d, received error: %v\n", i+1, table.expected, err)
		} else if err == nil && table.err != nil {
			t.Errorf("Test %d: expected error, received %d\n", i+1, actual)
		} else if err != nil && err.Error() != table.err.Error() {
			t.Errorf("Test %d: expected error %s, actual %s\n", i+1, table.err.Error(), err.Error())
		} else if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, table.expected, actual)
		}
	}
}
