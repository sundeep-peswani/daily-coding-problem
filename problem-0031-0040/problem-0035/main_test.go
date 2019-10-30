package main

import "testing"

func TestSegregateRGB(t *testing.T) {
	tests := []struct{
		input []byte
		expected []byte
	}{
		{[]byte{'G', 'B', 'R', 'R', 'B', 'R', 'G'}, []byte{'R', 'R', 'R', 'G', 'G', 'B', 'B'}},
	}

	for i, test := range tests {
		actual, changes := segregateRGB(test.input)
		if !isEqual(actual, test.expected) {
			t.Errorf("Test %d: expected %v, got %v\n", i+1, test.expected, actual)
		}
		if len(test.expected) < changes {
			t.Errorf("Test %d: made %d swaps, expected %d\n", i+1, changes, len(test.expected))
		}
	}
}

func isEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}