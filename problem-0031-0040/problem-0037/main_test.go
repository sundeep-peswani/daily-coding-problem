package main

import "testing"

func TestPowerSet(t *testing.T) {
	tests := []struct {
		set      []int
		expected [][]int
	}{
		{[]int{1}, [][]int{
			[]int{},
			[]int{1},
		}},
		{[]int{2}, [][]int{
			[]int{},
			[]int{2},
		}},
		{[]int{1, 2}, [][]int{
			[]int{},
			[]int{1},
			[]int{2},
			[]int{1, 2},
		}},
		{[]int{1, 2, 3}, [][]int{
			[]int{},
			[]int{1},
			[]int{2},
			[]int{3},
			[]int{1, 2},
			[]int{1, 3},
			[]int{2, 3},
			[]int{1, 2, 3},
		}},
	}

	for i, test := range tests {
		actual := powerSet(test.set)
		if !multiEqual(actual, test.expected) {
			t.Errorf("Test %d:\nExpected: %#v\nActual: %#v\n", i+1, test.expected, actual)
		}
	}
}

func multiEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, item := range b {
		if !multiHas(a, item) {
			return false
		}
	}

	return true
}

func multiHas(a [][]int, b []int) bool {
	for _, item := range a {
		if equal(item, b) {
			return true
		}
	}

	return false
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, i := range a {
		if !has(b, i) {
			return false
		}
	}

	return true
}

func has(a []int, b int) bool {
	for _, i := range a {
		if i == b {
			return true
		}
	}

	return false
}
