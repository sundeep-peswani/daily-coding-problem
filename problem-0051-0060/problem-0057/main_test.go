package main

import (
	"reflect"
	"testing"
)

func TestSplitIntoLines(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected []string
	}{
		{"abcdefghijklmnop", 1, nil},
		{"abcdefghijklmnop", 16, []string{"abcdefghijklmnop"}},
		{"abcdefghijklmnop", 17, []string{"abcdefghijklmnop"}},
		{"the quick brown fox jumps over the lazy dog", 10, []string{"the quick", "brown fox", "jumps over", "the lazy", "dog"}},
	}

	for i, test := range tests {
		actual := splitIntoLines(test.s, test.k)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Test %d: s = %s, k = %d\nexpected: %#v\nactual:   %#v\n", i+1, test.s, test.k, test.expected, actual)
		}
	}
}
