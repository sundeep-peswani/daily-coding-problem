package main

import "fmt"

func main() {

}

type node struct {
	val int
	left, right *node
}

func newNode(val int, left, right *node) *node {
	return &node{val, left, right}
}

func (n *node) isFull() bool {
	return (n.left == nil && n.right == nil) || (n.left != nil && n.right != nil)
}

func (n *node) convertToFull() *node {
	if n == nil {
		return nil
	}

	l, r := n.left.convertToFull(), n.right.convertToFull()
	if l == nil && r != nil {
		return r
	}

	if l != nil && r == nil {
		return l
	}

	return newNode(n.val, l, r)
}

func (n *node) equals(a *node) bool {
	if a == nil {
		return false
	}
	if n.val != a.val {
		return false
	}

	if n.left != nil {
		return n.left.equals(a.left)
	} else if a.left != nil {
		return false
	}

	if n.right != nil {
		return n.right.equals(a.right)
	} else if a.right != nil {
		return false
	}

	return true
}

func (n *node) String() string {
	if n == nil {
		return "()"
	}

	return fmt.Sprintf("(%d %s %s)", n.val, n.left, n.right)
}