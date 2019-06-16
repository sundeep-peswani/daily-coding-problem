package main

import "testing"

func TestRegexMatcher(t *testing.T) {
	tests := []struct {
		regex, given string
		expected     bool
	}{
		{"ray", "ray", true},
		{"ra.", "ray", true},
		{"", "ray", false},
		{".*", "", true},
		{"ra.", "raymond", false},
		{"ra.*", "raymond", true},
		{"ra.n", "rain", true},
		{".*at", "chat", true},
		{".*at", "chats", false},
		{"ra.*r", "racer", true},
		{"ra.*r", "razor", true},
		{"ra.*r", "randomizer", true},
		{".*ra.n.*", "raincoat", true},
		{".*ra.n.*", "drainage", true},
		{"a*.c*", "abc", true},
		{"a*.c*", "aaaaaaaabccccccc", true},
		{"a*.c*", "aaaaaaaabbccccccc", false},
		{"...", "abc", true},
		{".*", "anything", true},
		{".*.*.*g", "anything", true},
		{".*..g", "anything", true},
		{"bas.*ent", "basement", true},
	}

	for i, test := range tests {
		actual := regexMatch(test.regex, test.given)
		if actual != test.expected {
			t.Errorf("Test %d: /%s/.match(%s) - expected %t, actual %t\n", i+1, test.regex, test.given, test.expected, actual)
		}
	}
}
