package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestReconstruct(t *testing.T) {
	tables := []struct {
		dictionary []string
		given      string
		expected   [][]string
	}{
		{[]string{"fox", "quick", "the", "brown"}, "thequickbrownfox", [][]string{strings.Split("the quick brown fox", " ")}},
		{[]string{"bed", "bath", "bedbath", "beyond", "and"}, "bedbathandbeyond", [][]string{strings.Split("bed bath and beyond", " "), strings.Split("bedbath and beyond", " ")}},
		{[]string{"bed", "bedbath", "beyond", "and"}, "bedbathandbeyond", [][]string{strings.Split("bedbath and beyond", " ")}},
		{[]string{"bed", "beyond", "and"}, "bedbathandbeyond", [][]string{}},
		{[]string{"bed", "bath", "bedbath", "beyonder", "and"}, "bedbathandbeyond", [][]string{}},
	}

	for i, table := range tables {
		actual := reconstruct(table.dictionary, table.given)

		fmt.Println(actual)

		found := len(table.expected) == 0
		for _, ex := range table.expected {
			found = found || reflect.DeepEqual(actual, ex)
		}

		if !found {
			t.Errorf("Test %d: unable to found %s in %v\n", i+1, table.given, table.dictionary)
		}
	}
}
