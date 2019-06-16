package main

import (
	"fmt"
)

// A Node of a binary tree
type Node struct {
	val   int
	left  *Node
	right *Node
}

func main() {
	root := Node{
		0,
		&Node{
			1,
			nil,
			nil,
		},
		&Node{
			0,
			&Node{
				1,
				&Node{1, nil, nil},
				&Node{1, nil, nil},
			},
			&Node{0, nil, nil},
		},
	}

	count, _ := root.countUnival()
	fmt.Println(count)
}

func (n *Node) countUnival() (int, bool) {
	if n == nil {
		return 0, false
	}

	if n.left == nil && n.right == nil {
		return 1, true
	}

	lTotal, lIsUnival := n.left.countUnival()
	rTotal, rIsUnival := n.right.countUnival()

	if lIsUnival && n.left.val == n.val && rIsUnival && n.right.val == n.val {
		return 1 + lTotal + rTotal, true
	}

	return lTotal + rTotal, false
}
