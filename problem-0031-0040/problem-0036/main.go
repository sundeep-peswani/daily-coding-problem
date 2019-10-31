package main

func main() {

}

func findSecondLargest(root *node) int {
	if root.right == nil {
		if root.left == nil {
			return -1
		}

		return root.left.val
	}

	second := findSecondLargest(root.right)
	if second == -1 {
		return root.val
	}

	return second
}

type node struct {
	val int
	left, right *node
}

func createBST(input []int) *node {
	if len(input) == 0 {
		return nil
	}

	if len(input) == 1 {
		return &node{input[0], nil, nil}
	}

	middle := len(input) / 2
	left, right := createBST(input[:middle]), createBST(input[middle+1:])
	return &node{input[middle], left, right}
}
