package main

import (
	"fmt"
	"log"
	"strconv"
)

// A Node in an XOR Linked List
type Node struct {
	value string
	heap  *Heap
	both  int64
}

// A Heap is a representation of memory so we can
// dereference pointers in go
type Heap struct {
	memory map[int64]*Node
}

func main() {
	h := newHeap()

	head, err := h.newNode("first", 0)
	if err != nil {
		log.Fatalf("Unable to create linked list: %v\n", err)
	}

	for _, value := range []string{"second", "third", "fourth"} {
		if err := head.add(value); err != nil {
			log.Fatalf("Unable to add %s: %v\n", value, err)
		}
	}

	if n, err := head.get(2); err != nil {
		log.Fatalf("Error encountered: %v\n", err)
	} else {
		fmt.Printf("Item at index 2 is %s\n", n.value)
	}
}

func newHeap() Heap {
	h := Heap{}
	h.memory = make(map[int64]*Node)

	return h
}

func (h *Heap) newNode(value string, prevAddr int64) (*Node, error) {
	n := Node{value, h, prevAddr}

	var p int64
	var err error

	if p, err = n.getPointer(); err != nil {
		return nil, err
	}

	h.memory[p] = &n
	return &n, nil
}

func (n *Node) add(element string) error {
	curr := n
	addr := int64(0)

	var next *Node
	var err error

	// forward to the end
	for {
		if next, err = curr.getNext(addr); err != nil {
			return err
		} else if next == nil {
			break
		} else {
			if addr, err = curr.getPointer(); err != nil {
				return err
			}
			curr = next
		}
	}

	if addr, err = curr.getPointer(); err != nil {
		return err
	}

	if n, err := curr.heap.newNode(element, addr); err != nil {
		return err
	} else if addr, err = n.getPointer(); err != nil {
		return err
	}

	curr.both = curr.both ^ addr

	return nil
}

func (n *Node) get(index int) (*Node, error) {
	if index == 0 {
		return n, nil
	}

	var addr int64
	var err error
	var curr, next *Node

	curr = n

	for i := 1; i <= index; i++ {
		if next, err = curr.getNext(addr); err != nil {
			return nil, err
		} else if addr, err = curr.getPointer(); err != nil {
			return nil, err
		}
		curr = next
	}

	return curr, nil
}

func (n Node) getNext(currentAddr int64) (*Node, error) {
	p := n.both ^ currentAddr
	if p == 0 {
		return nil, nil
	}

	return n.heap.dereferencePointer(p)
}

func (n *Node) getPointer() (int64, error) {
	return strconv.ParseInt(fmt.Sprintf("%p", n), 0, 0)
}

func (h *Heap) dereferencePointer(addr int64) (*Node, error) {
	if n, ok := h.memory[addr]; ok {
		return n, nil
	}

	return nil, fmt.Errorf("could not find node at addr: %d", addr)
}
