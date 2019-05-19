package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestFindIntersection(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tables := []struct {
		lenIntersection  int
		intersectingNode *node
	}{
		{0, nil},
		{1, newNode(rand.Intn(1000))},
		{rand.Intn(100), newNode(rand.Intn(1000))},
	}

	for i, table := range tables {
		lenA, lenB := 100+rand.Intn(100), 100+rand.Intn(100)
		a, b := newLinkedList(), newLinkedList()

		for i := 0; i < lenA-table.lenIntersection; i++ {
			a.push(newNode(rand.Intn(1000)))
		}

		for i := 0; i < lenB-table.lenIntersection; i++ {
			b.push(newNode(rand.Intn(1000)))
		}

		a.push(table.intersectingNode)
		b.push(table.intersectingNode)

		for i := 0; i < table.lenIntersection-1; i++ {
			n := newNode(rand.Intn(1000))
			a.push(n)
			b.push(n)
		}

		actual := findIntersection(a, b)
		if actual != table.intersectingNode {
			t.Errorf("Test %d: failed to find correct intersection node, expected '%s', actual '%s': \n", i+1, table.intersectingNode, actual)
			t.Errorf("A: %s\nB: %s\n", a, b)
		}
	}
}
