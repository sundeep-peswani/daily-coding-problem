package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{9, 11, 8, 5, 7, 10}, 5},
	}

	for i, test := range tests {
		actual := maxProfit(test.prices)
		if test.expected != actual {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, test.expected, actual)
		}
	}
}
