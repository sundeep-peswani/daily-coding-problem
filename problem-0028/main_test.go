package main

import (
	"strings"
	"testing"
)

func TestJustifyText(t *testing.T) {
	tests := []struct {
		words    []string
		length   int
		expected []string
	}{
		{
			strings.Split("the quick brown fox jumps over the lazy dog", " "),
			16,
			strings.Split("the  quick brown\nfox  jumps  over\nthe   lazy   dog", "\n"),
		},
	}

	for i, test := range tests {
		actual := justifyText(test.words, test.length)
		if !isEqual(actual, test.expected) {
			t.Errorf("Test %d: failed to justify text.\nExpected: %#v\nActual  : %#v\n", i+1, test.expected, actual)
		}
	}
}

func isEqual(a, b []string) bool {
	for i, line := range a {
		if b[i] != line {
			return false
		}
	}

	return true
}
