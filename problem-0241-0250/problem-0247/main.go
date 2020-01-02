package main

func main() {

}

type node struct {
	val int
	left, right *node
}

func newNode(val int, left, right *node) *node {
	return &node{val, left, right}
}

func (n node) height() int {
	return max(n.leftHeight(), n.rightHeight()) + 1
}

func (n node) leftHeight() int {
	if n.left == nil {
		return 0
	}

	return n.left.height()
}

func (n node) rightHeight() int {
	if n.right == nil {
		return 0
	}

	return n.right.height()
}

func (n node) isHeightBalanced() bool {
	return abs(n.leftHeight() - n.rightHeight()) <= 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	
	return n
}