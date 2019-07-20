package main

func main() {

}

type node struct {
	value       int
	left, right *node
}

func findDeepestNode(root *node) *node {
	if root == nil {
		return nil
	}

	var last *node

	q := []*node{root}
	for i := 0; i < len(q); i++ {
		last = q[i]

		if last.left != nil {
			q = append(q, last.left)
		}
		if last.right != nil {
			q = append(q, last.right)
		}
	}

	return last
}
