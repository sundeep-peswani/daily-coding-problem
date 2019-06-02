package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestLinkedListRemove(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 10; i < 100000000; i *= 10 {
		arr := randArray(i)
		k := rand.Intn(i)

		arr[i-k] = i + 1
		ll := newLinkedListFrom(arr)

		if !ll.remove(k) {
			t.Errorf("Test with %d elements: unable to remove %dth item\n", i, k)
		}

		n := ll.head
		for j := 0; j < i-k; j++ {
			if n.value >= i {
				t.Errorf("Test with %d elements: found unexpected item %d at %d, after removing %dth-last item\n", i, n.value, j+1, k)
			}
			n = n.next
		}
	}
}

func randArray(i int) []int {
	res := make([]int, i)
	for j := 0; j < i; j++ {
		res[j] = rand.Intn(i)
	}
	return res
}

func newLinkedListFrom(arr []int) *linkedList {
	ll := newLinkedList()

	for _, n := range arr {
		ll.pushBack(n)
	}

	return ll
}
