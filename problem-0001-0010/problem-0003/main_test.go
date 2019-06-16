package main

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	tables := []struct {
		str, val, left, right string
		err                   error
	}{
		{"{,,}", "", "", "", nil},
		{"{root,,}", "root", "", "", nil},
		{"{root,{left,,},{right,,}}", "root", "{left,,}", "{right,,}", nil},
		{"{root,{left,{left.left,,{left.left.right}},},{right,,}}", "root", "{left,{left.left,,{left.left.right}},}", "{right,,}", nil},
		{"{root,,", "", "", "", fmt.Errorf("invalid input, non-tree structure: {root,,")},
		{"root,,}", "", "", "", fmt.Errorf("invalid input, non-tree structure: root,,}")},
		{"{root|left|right}", "", "", "", fmt.Errorf("invalid input, no comma found: {root|left|right}")},
		{"{root,{{{{{{{{}}}}", "", "", "", fmt.Errorf("invalid input: {root,{{{{{{{{}}}}")},
	}

	for i, table := range tables {
		val, left, right, err := split(table.str)
		if err != nil && table.err == nil {
			t.Errorf("Test %d: unexpected error: %v", i, err)
		}
		if err == nil && table.err != nil {
			t.Errorf("Test %d: expected error, got none", i)
		}
		if err != nil && table.err != nil && err.Error() != table.err.Error() {
			t.Errorf("Test %d: error - expected '%s', actual '%s'", i, table.err.Error(), err.Error())
		}
		if val != table.val {
			t.Errorf("Test %d: val - expected '%s', actual '%s'", i, table.val, val)
		}
		if left != table.left {
			t.Errorf("Test %d: left - expected '%s', actual '%s'", i, table.left, left)
		}
		if right != table.right {
			t.Errorf("Test %d: right - expected '%s', actual '%s'", i, table.right, right)
		}
	}
}

func TestSerialize(t *testing.T) {
	tables := []struct {
		root     Node
		expected string
	}{
		{Node{"root", nil, nil}, "{root,,}"},
		{Node{"root", &Node{"left", nil, nil}, nil}, "{root,{left,,},}"},
		{Node{"root", &Node{"left", nil, nil}, &Node{"right", nil, nil}}, "{root,{left,,},{right,,}}"},
		{Node{"root", &Node{"left", &Node{"left.left", nil, nil}, nil}, &Node{"right", nil, &Node{"right.right", nil, nil}}}, "{root,{left,{left.left,,},},{right,,{right.right,,}}}"},
	}

	for i, table := range tables {
		actual := serialize(&table.root)
		if actual != table.expected {
			t.Errorf("Test %d: expected \"%s\", actual \"%s\"", i, table.expected, actual)
		}
	}
}

func TestDeserialize(t *testing.T) {
	tables := []struct {
		str      string
		expected Node
	}{
		{"{root,,}", Node{"root", nil, nil}},
		{"{root,{left,,},}", Node{"root", &Node{"left", nil, nil}, nil}},
		{"{root,{left,,},{right,,}}", Node{"root", &Node{"left", nil, nil}, &Node{"right", nil, nil}}},
		{"{root,{left,{left.left,,},},{right,,{right.right,,}}}", Node{"root", &Node{"left", &Node{"left.left", nil, nil}, nil}, &Node{"right", nil, &Node{"right.right", nil, nil}}}},
	}

	for i, table := range tables {
		actual, err := deserialize(table.str)
		if err != nil {
			t.Errorf("Test %d: error %v", i, err)
		}
		if !equalTree(&table.expected, actual) {
			t.Errorf("Test %d: expected %v, actual %v", i, table.expected, actual)
		}
	}
}

func TestSerializeAndDeserialize(t *testing.T) {
	tables := []struct {
		expected Node
	}{
		{Node{"root", nil, nil}},
		{Node{"root", &Node{"left", nil, nil}, nil}},
		{Node{"root", &Node{"left", nil, nil}, &Node{"right", nil, nil}}},
		{Node{"root", &Node{"left", &Node{"left.left", nil, nil}, nil}, &Node{"right", nil, &Node{"right.right", nil, nil}}}},
	}

	for i, table := range tables {
		actual, err := deserialize(serialize(&table.expected))
		if err != nil {
			t.Errorf("Test %d: error %v", i, err)
		}
		if !equalTree(&table.expected, actual) {
			t.Errorf("Test %d: expected %v, actual %v", i, table.expected, actual)
		}
	}
}

func equalTree(a, b *Node) bool {
	if a == nil {
		return b == nil
	} else if b == nil {
		return a == nil
	}

	if a.val != b.val {
		return false
	}

	return equalTree(a.left, b.left) && equalTree(a.right, b.right)
}
