package main

type node struct {
	value int
	parent, left, right *node
}

func nextBiggest(root *node, search int) int {
	n := findBST(root, search)

	if n.right != nil {
		// return leftmost
		n = n.right
		for n.left != nil {
			n = n.left
		}
		return n.value
	}

	if n.parent == nil {
		return search
	}

	if n.parent.left == n {
		return n.parent.value
	}

	// n is the right node and is therefore bigger than n.parent
	// however, it is *smaller* than n's grandparent
	return n.parent.parent.value
}

func findBST(root *node, search int) *node {
	if root == nil {
		return nil
	}

	if root.value == search {
		return root
	}

	if root.value > search {
		return findBST(root.left, search)
	}

	return findBST(root.right, search)
}
