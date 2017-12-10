package main

import "fmt"

func main() {
	n5 := &TreeNode{5, nil, nil}
	n4 := &TreeNode{4, n5, nil}
	n3 := &TreeNode{3, n4, nil}
	// n2 := &TreeNode{2, nil, nil}
	// n1 := &TreeNode{1, nil, n2}
	n0 := &TreeNode{0, n3, nil}
	fmt.Println(isBalanced(n0))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return getHighestHeight(root, 0) <= getLowestHeight(root, 0)+1
}

func getLowestHeight(n *TreeNode, h int) int {
	if n.Left == nil || n.Right == nil {
		return h + 1
	}

	lh := getLowestHeight(n.Left, h+1)
	rh := getLowestHeight(n.Right, h+1)
	if lh > rh {
		return rh
	}
	return lh
}

func getHighestHeight(n *TreeNode, h int) int {
	if n.Left == nil && n.Right == nil {
		return h + 1
	} else if n.Left == nil && n.Right != nil {
		return getHighestHeight(n.Right, h+1)
	} else if n.Right == nil && n.Left != nil {
		return getHighestHeight(n.Left, h+1)
	}

	lh := getHighestHeight(n.Left, h+1)
	rh := getHighestHeight(n.Right, h+1)
	if lh > rh {
		return lh
	}
	return rh
}
