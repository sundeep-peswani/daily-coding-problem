package main

import (
	"testing"
)

func TestLocking(t *testing.T) {
	root := b(1, l(2, z(3)), r(4, z(5)))

	if !root.left.lock() {
		t.Errorf("Unable to lock root.left\n")
	}

	if !root.left.isLocked() {
		t.Errorf("root.left should be locked\n")
	}

	if !root.left.left.lock() {
		t.Errorf("Unable to lock root.left.left\n")
	}

	if !root.left.isLocked() {
		t.Errorf("root.left.left should be locked\n")
	}

	if root.lock() {
		t.Errorf("root is unexpectedly unlockable\n")
	}

	if root.right.unlock() {
		t.Errorf("root.right is unexpectedly unlockable\n")
	}

	if root.left.left.unlock() {
		t.Errorf("Able to unlock root.left.left with locked root.left\n")
	}

	if !root.left.left.isLocked() {
		t.Errorf("root.left.left is unexpectedly unlocked\n")
	}

	if !root.left.unlock() {
		t.Errorf("Unable to unlock root.left\n")
	}

	if root.left.isLocked() {
		t.Errorf("root.left is still locked\n")
	}

	if !root.left.left.unlock() {
		t.Errorf("Unable to unlock root.left.left\n")
	}

	if root.left.left.isLocked() {
		t.Errorf("root.left.left is still locked\n")
	}
}

func z(val int) *node {
	return newNode(val)
}

func l(val int, left *node) *node {
	n := z(val)
	n.left = left
	left.parent = n

	return n
}

func r(val int, right *node) *node {
	n := z(val)
	n.right = right
	right.parent = n

	return n
}

func b(val int, left, right *node) *node {
	n := l(val, left)
	n.right = right
	right.parent = n

	return n
}
