package main

import "testing"

func TestIsMatching(t *testing.T) {
	tests := []struct {
		given    string
		expected bool
	}{
		{"([])[]({})", true},
		{"([)]", false},
		{"((()", false},
		{"{", false},
		{")(", false},
		{"((()))[[[]]]{{{{}}}}", true},
		{"", true},
	}

	for i, test := range tests {
		actual := isMatching(test.given)
		if actual != test.expected {
			t.Errorf("Test %d: for '%s', expected %t, actual %t\n", i+1, test.given, test.expected, actual)
		}
	}
}
