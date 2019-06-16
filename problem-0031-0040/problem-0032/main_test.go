package main

import (
	"testing"
)

func TestFindArbitrage(t *testing.T) {
	tests := []struct {
		rates    [][]float64
		expected bool
	}{
		{
			[][]float64{
				[]float64{1, 0.5, 0.25},
				[]float64{2, 1, 0.5},
				[]float64{4, 0.25, 1},
			},
			false,
		},
		{
			[][]float64{
				[]float64{1, 2, 3, 4},
				[]float64{2, 1, 1, 1},
				[]float64{3, 1, 1, 1},
				[]float64{4, 1, 1, 1},
			},
			true,
		},
		{
			[][]float64{
				[]float64{1, 2, 4, 8},
				[]float64{0.5, 1, 2, 4},
				[]float64{0.25, 0.5, 1, 2},
				[]float64{0.125, 0.25, 0.5, 1},
			},
			false,
		},
	}

	for i, test := range tests {
		actual := findArbitrage(test.rates)
		if actual != test.expected {
			t.Errorf("Test %d: expected %v, actual %v\n", i+1, test.expected, actual)
		}
	}
}
