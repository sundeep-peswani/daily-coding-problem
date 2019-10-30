package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculateMedian(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []float64
	}{
		{[]int{2, 1, 5, 7, 2, 0, 5}, []float64{2, 1.5, 2, 3.5, 2, 2, 2}},
	}

	for i, test := range tests {
		var actual []float64
		lower, higher := newHeap(greaterThan), newHeap(lessThan)

		for _, n := range test.nums {
			actual = append(actual, calculateMedian(lower, higher, n))
		}

		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Test %d: expected %v, actual %v\n", i+1, test.expected, actual)
		}
	}
}

func TestHeap(t *testing.T) {
	tests := []struct{
		input []int
		cmp func(a, b int) bool
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, greaterThan, []int{6, 5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6}, lessThan, []int{1, 2, 3, 4, 5, 6}},
	}

	for i, test := range tests {
		h := newHeap(test.cmp)

		for _, v := range test.input {
			h.insert(v)
			fmt.Println(h)
		}

		fmt.Println(h)

		if h.get() != test.expected[0] {
			t.Errorf("Test %d: failed to get top of heap: %d\n", i + 1, h.get())
		}

		for j := 0; j < len(test.expected); j++ {
			removed := h.remove()
			if removed != test.expected[j] {
				t.Errorf("Test %d: removed %d, expected %d\n", i + 1, removed, test.expected[j])
			}
			fmt.Println(h)
		}

		if !h.isEmpty() {
			t.Errorf("Test %d: expected heap to be empty, has %d items left\n", i+1, h.length())
		}
	}
}
