package main

import "fmt"
import "math"

func main() {
	n5 := &TreeNode{5, nil, nil}
	n4 := &TreeNode{4, n5, nil}
	n3 := &TreeNode{3, n4, nil}
	// n2 := &TreeNode{2, nil, nil}
	// n1 := &TreeNode{1, nil, n2}
	n0 := &TreeNode{0, n3, nil}
	fmt.Println(isBalanced(n0))
	fmt.Println(isBalanced2(n0))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if depth(root) == 0 {
		return true
	}

	l := depth(root.Left)
	r := depth(root.Right)

	return math.Abs(float64(l-r)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := depth(node.Left)
	r := depth(node.Right)
	if l > r {
		return l + 1
	}

	return r + 1
}

var inbalanced = -100

func isBalanced2(root *TreeNode) bool {
	return tellBalance(root) != inbalanced
}

func tellBalance(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := tellBalance(node.Left)
	r := tellBalance(node.Right)
	if l == inbalanced || r == inbalanced || math.Abs(float64(l-r)) > 1 {
		return inbalanced
	}

	if l > r {
		return 1 + l
	}
	return 1 + r
}
