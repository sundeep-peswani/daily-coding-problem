package main

import "testing"

func Test_Zigzag(t *testing.T) {
	tests := []struct{
		input string
		k int
		expected []string
	}{
		{"thisisazigzag", 4, []string{
			"t     a     g",
			" h   s z   a",
			"  i i   i z",
			"   s     g",
		}},
	}

	for i, test := range tests {
		actual := zigzag(test.input, test.k)
		if !equal(test.expected, actual) {
			t.Errorf("Test %d: input = %s, k = %d\nexpected:\n%s\nactual:\n%s\n", i+1, test.input, test.k, test.expected, actual)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, row := range a {
		if row != b[i] {
			return false
		}
	}

	return true
}