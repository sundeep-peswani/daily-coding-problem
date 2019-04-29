package main

import (
	"fmt"
	"os"
	"strings"
)

// A Node of a binary tree
type Node struct {
	val   string
	left  *Node
	right *Node
}

func main() {
	root := Node{"root", &Node{"left", &Node{"left.left", nil, nil}, nil}, &Node{"right", nil, nil}}
	result, err := deserialize(serialize(&root))

	if err != nil {
		fmt.Printf("Error in deserializing: %v\n", err)
		os.Exit(1)
	}

	if result.left.left.val != "left.left" {
		fmt.Printf("Failed to correctly marshal %v", root)
		os.Exit(1)
	}
}

func serialize(n *Node) string {
	if n == nil {
		return ""
	}

	return fmt.Sprintf("{%s,%s,%s}", n.val, serialize(n.left), serialize(n.right))
}

func deserialize(str string) (*Node, error) {
	if len(str) == 0 {
		return nil, nil
	}

	val, left, right, err := split(str)
	if err != nil {
		return nil, err
	}

	root := Node{val, nil, nil}

	if leftNode, err := deserialize(left); err != nil {
		return nil, err
	} else {
		root.left = leftNode
	}

	if rightNode, err := deserialize(right); err != nil {
		return nil, err
	} else {
		root.right = rightNode
	}

	return &root, nil
}

func split(str string) (string, string, string, error) {
	if str[0] != '{' || str[len(str)-1] != '}' {
		return "", "", "", fmt.Errorf("invalid input, non-tree structure: %s", str)
	}

	i := strings.Index(str, ",")
	if i == -1 {
		return "", "", "", fmt.Errorf("invalid input, no comma found: %s", str)
	}

	val, rest := str[1:i], str[i+1:len(str)-1]

	fmt.Printf("from '%s', val = '%s', rest = '%s'\n", str, val, rest)

	braces := 0
	for i := 0; i < len(rest); i++ {
		if rest[i] == '{' {
			braces++
		} else if rest[i] == '}' {
			braces--
		} else if rest[i] == ',' && braces == 0 {
			return val, rest[:i], rest[i+1:], nil
		}
	}

	return "", "", "", fmt.Errorf("invalid input: %s", str)
}
