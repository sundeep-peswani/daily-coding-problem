package main

import (
	"testing"
)

func TestIsMappable(t *testing.T) {
	tests := []struct {
		s1, s2 string
		expected bool
	}{
		{"abc", "bcd", true},
		{"foo", "bar", false},
		{"yes", "yes", true},
		{"eyes", "yes", false},
		{"noon", "yaya", true},
		{"aaabbb", "bababa", true},
		{"aaabbb", "nonono", true},
	}

	for i, test := range tests {
		if isMappable(test.s1, test.s2) != test.expected {
			t.Errorf("Test %d: %s, %s, expected %v, got %v\n", i+1, test.s1, test.s2, test.expected, !test.expected)
		}
	}
}