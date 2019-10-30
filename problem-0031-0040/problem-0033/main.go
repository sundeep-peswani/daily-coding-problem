package main

import (
	"fmt"
	"os"
)

func main() {

}

func usage() {
	fmt.Printf("Usage: %s\n[ints...]\n", os.Args[0])
	os.Exit(1)
}

func calculateMedian(lower, higher *heap, n int) float64 {
	if lower.isEmpty() && higher.isEmpty() {
		lower.insert(n)
		return float64(n)
	}

	median := getMedian(lower, higher)
	diff := lower.length() - higher.length()
	isLowerThanMedian := float64(n) < median

	switch (diff) {
	case 0:
		if isLowerThanMedian {
			lower.insert(n)
		} else {
			higher.insert(n)
		}
	case 1:
		if isLowerThanMedian {
			higher.insert(lower.remove())
			lower.insert(n)
		} else {
			higher.insert(n)
		}
	case -1:
		if isLowerThanMedian {
			lower.insert(n)
		} else {
			lower.insert(higher.remove())
			higher.insert(n)
		}
	}

	return getMedian(lower, higher)
}

func getMedian(lower, higher *heap) float64 {
	if lower.length() == higher.length() {
		return float64(lower.get() + higher.get()) / 2
	}

	if lower.length() > higher.length() {
		return float64(lower.get())
	}

	return float64(higher.get())
}

type heap struct {
	// return true if a should be placed 'higher' in the heap
	cmp		func(a, b int) bool
	data	[]int
}

func greaterThan(a, b int) bool { return a > b }
func lessThan(a, b int) bool { return a < b }

func newHeap(cmp func (a, b int) bool) *heap {
	return &heap{cmp, make([]int, 0)}
}

func (h *heap) get() int {
	if len(h.data) == 0 {
		return -1
	}

	return h.data[0]
}

func (h *heap) insert(n int) {
	if h.isEmpty() {
		h.data = append(h.data, n)
		return
	}
	
	h.data = append(h.data, n)
	
	idx := h.length()
	parent := h.parent(idx)
	for idx > 0 && h.cmp(h.data[idx], h.data[parent]) {
		h.data[idx], h.data[parent] = h.data[parent], h.data[idx]
		idx = parent
		parent = h.parent(idx)
	}
}

func (h *heap) remove() int {
	if h.isEmpty() {
		return -1
	}

	n := h.get()
	h.data[0], h.data[h.length()] = h.data[h.length()], h.data[0]
	h.data = h.data[:h.length()]

	idx := 0
	left, right := h.left(idx), h.right(idx)
	child := right

	for idx <= h.length() {
		if idx < h.length() {
			if right > h.length() || h.cmp(h.data[left], h.data[right]) {
				child = left
			}
		}

		if child > h.length() || h.cmp(h.data[idx], h.data[child]) {
			return n
		}

		h.data[idx], h.data[child] = h.data[child], h.data[idx]
		idx = child
		left, right = h.left(idx), h.right(idx)
		child = right
	}

	return n
}

func (h *heap) length() int {
	return len(h.data) - 1
}

func (h *heap) isEmpty() bool {
	return len(h.data) == 0
}

func (h *heap) left(idx int) int {
	return 2 * idx + 1;
}

func (h *heap) right(idx int) int {
	return 2 * (idx + 1);
}

func (h *heap) parent(idx int) int {
	if (idx <= 0) {
		return -1
	}

	return (idx - 1) / 2;
}

func (h *heap) String() string {
	return fmt.Sprintf("%#v", h.data)
}