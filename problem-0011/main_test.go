package main

import (
	"math/rand"
	"testing"
)

func TestDictionary(t *testing.T) {
	testWords := []string{
		"",
		"atest",
		"test",
		"testes",
		"testing",
		"testy",
	}
	invalidWords := []string{
		"btest",
		"c",
		"chucktesta",
		"code",
		"tessa",
		"tesseract",
		"testachuck",
		"tester",
	}

	dictionary := buildDictionary(testWords)

	for i, word := range testWords {
		if len(word) == 0 {
			continue
		}

		prefixLength := max(1, rand.Intn(len(word)))
		prefix := word[0:prefixLength]

		actualWords := dictionary.find(prefix)
		if !contains(actualWords, word) {
			t.Errorf("Valid test %d: expected to find %s with prefix %s, not found in %v\n", i, word, prefix, actualWords)
		}
	}

	for i, word := range invalidWords {
		prefixLength := max(1, rand.Intn(len(word)))
		prefix := word[0:prefixLength]

		actualWords := dictionary.find(prefix)
		if contains(actualWords, word) {
			t.Errorf("Invalid test %d: found unexpected word %s, with prefix %s, in list: %v\n", i, word, prefix, actualWords)
		}
	}
}

func contains(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
