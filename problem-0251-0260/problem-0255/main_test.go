package main

import "testing"

func Test_CalculateTransitiveClosure(t *testing.T) {
	tests := []struct{
		graph [][]int
		expected [][]int
	}{
		{[][]int{
			[]int{0, 1, 3},
			[]int{1, 2},
			[]int{2},
			[]int{3},
		}, [][]int{
			[]int{1, 1, 1, 1},
			[]int{0, 1, 1, 0},
			[]int{0, 0, 1, 0},
			[]int{0, 0, 0, 1},
		}},
		{[][]int{
			[]int{0},
			[]int{1},
			[]int{2},
			[]int{3},
		}, [][]int{
			[]int{1, 0, 0, 0},
			[]int{0, 1, 0, 0},
			[]int{0, 0, 1, 0},
			[]int{0, 0, 0, 1},
		}},
		{[][]int{
			[]int{0, 1},
			[]int{1, 2},
			[]int{2, 3},
			[]int{3, 0},
		}, [][]int{
			[]int{1, 1, 1, 1},
			[]int{1, 1, 1, 1},
			[]int{1, 1, 1, 1},
			[]int{1, 1, 1, 1},
		}},		
	}

	for i, test := range tests {
		actual := calculateTransitiveClosure(test.graph)
		if !equal(test.expected, actual) {
			t.Errorf("Test %d:\nexpected: %#v\nactual:   %#v\n", i+1, test.expected, actual)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, row := range a {
		if len(row) != len(b[i]) {
			return false
		}

		for j, item := range row {
			if item != b[i][j] {
				return false
			}
		}
	}

	return true
}
