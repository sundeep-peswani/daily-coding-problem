package main

import (
	"fmt"
)

func main() {

}

type node struct {
	value               int
	locked              bool
	left, right, parent *node
}

func newNode(value int) *node {
	return &node{value, false, nil, nil, nil}
}

func (n *node) String() string {
	if n == nil {
		return "x"
	}

	return fmt.Sprintf("(%s %d %s)", n.left, n.value, n.right)
}

func (n *node) lock() bool {
	if !n.isLockable() {
		return false
	}
	n.locked = true
	return true
}

func (n *node) unlock() bool {
	if !n.isUnlockable() {
		return false
	}

	n.locked = false
	return true
}

func (n *node) isLocked() bool {
	return n.locked
}

func (n *node) isLockable() bool {
	if n == nil {
		return true
	}

	return !n.isLocked() && n.left.isLockable() && n.right.isLockable()
}

func (n *node) isUnlockable() bool {
	if n == nil || !n.isLocked() {
		return false
	}

	c := n.parent
	for c != nil {
		if c.isLocked() {
			return false
		}
		c = c.parent
	}

	return true
}
