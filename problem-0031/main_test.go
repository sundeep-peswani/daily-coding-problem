package main

import "testing"

func TestEditDistance(t *testing.T) {
	tests := []struct {
		a, b     string
		expected int
	}{
		{"kitten", "sitting", 3},
	}

	for i, test := range tests {
		actual := editDistance(test.a, test.b)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, test.expected, actual)
		}
	}
}
