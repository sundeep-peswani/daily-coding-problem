package main

import "strconv"

func main() {

}

type node struct {
	value       string
	result      int
	left, right *node
}

func newNode(value string) *node {
	if value == "+" || value == "-" || value == "*" || value == "/" {
		return &node{value, 0, nil, nil}
	}

	if n, err := strconv.Atoi(value); err == nil {
		return &node{value, n, nil, nil}
	}

	return nil
}

func (n *node) evaluate() int {
	l, r := 0, 0

	if n.left != nil {
		l = n.left.evaluate()
	}
	if n.right != nil {
		r = n.right.evaluate()
	}

	switch n.value {
	case "+":
		n.result = l + r
	case "-":
		n.result = l - r
	case "*":
		n.result = l * r
	case "/":
		n.result = l / r
	}

	return n.result
}
