package main

func main() {

}

type binaryNode struct {
	value int
	left, right *binaryNode
}

func (root *binaryNode) find(k int) *binaryNode {
	if k == root.value {
		return root
	}

	if k < root.value && root.left != nil {
		return root.left.find(k)
	}

	if k > root.value && root.right != nil {
		return root.right.find(k)
	}

	return nil
}

func findSum(root *binaryNode, k int) []*binaryNode {
	q := []*binaryNode{root}

	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}

		pair := root.find(k - n.value)
		if pair != nil && pair != n {
			return []*binaryNode{n, pair}
		}
	}

	return nil
}