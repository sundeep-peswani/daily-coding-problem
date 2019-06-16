package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}

	var arr []int
	for i := 2; i < len(os.Args); i++ {
		n, err := strconv.Atoi(os.Args[i])
		if err != nil {
			usage()
		}
		arr = append(arr, n)
	}

	fmt.Println(maxSubarray(arr, k))
}

func usage() {
	fmt.Printf("Usage: %s <k> <...array of integers>\n", os.Args[0])
	os.Exit(1)
}

func maxSubarray(arr []int, k int) []int {
	n := len(arr)
	if k == 0 || n == 0 {
		return []int{}
	}

	d := &deque{nil, nil}
	var res []int

	// pick the largest in the
	for i := 0; i < k; i++ {
		for !d.isEmpty() && arr[d.peekBack().val] < arr[i] {
			d.popBack()
		}
		d.pushBack(&node{i, nil, nil})
	}

	for i := k; i < n; i++ {
		res = append(res, arr[d.peek().val])
		for !d.isEmpty() && arr[d.peekBack().val] < arr[i] {
			d.popBack()
		}

		for !d.isEmpty() && d.peek().val <= i-k {
			d.pop()
		}

		d.pushBack(&node{i, nil, nil})
	}

	res = append(res, arr[d.pop().val])

	return res
}

type node struct {
	val        int
	next, prev *node
}

type deque struct {
	head, tail *node
}

func (d *deque) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Head: %#v, Tail: %#v, Queue: ", &d.head, &d.tail))
	n := d.head
	for n != nil {
		sb.WriteString(fmt.Sprintf("%d ", n.val))
		n = n.next
	}

	return sb.String()
}

func (d *deque) isEmpty() bool {
	return d.head == nil && d.tail == nil
}

func (d *deque) push(n *node) {
	n.prev = nil
	n.next = d.head

	if d.head != nil {
		d.head.prev = n
	}
	d.head = n

	if d.tail == nil {
		d.tail = n
	}
}

func (d *deque) pop() *node {
	n := d.head
	if n != nil {
		d.head = n.next
		n.prev = nil
		n.next = nil

		if d.head != nil {
			d.head.prev = nil
		}

		if d.tail == n {
			d.popBack()
		}
	}
	return n
}

func (d *deque) pushBack(n *node) {
	if d.tail != nil {
		d.tail.next = n
	}
	n.prev = d.tail
	n.next = nil
	d.tail = n

	if d.head == nil {
		d.head = n
	}
}

func (d *deque) popBack() *node {
	n := d.tail
	if n != nil {
		d.tail = n.prev
		n.prev = nil
		n.next = nil

		if d.tail != nil {
			d.tail.next = nil
		}

		if d.head == n {
			d.pop()
		}
	}
	return n
}

func (d *deque) peek() *node {
	return d.head
}

func (d *deque) peekBack() *node {
	return d.tail
}
