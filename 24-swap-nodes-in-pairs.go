package main

import "fmt"

func main() {
	n5 := &ListNode{5, nil}
	swapPairs(n5).status()

	n4 := &ListNode{4, n5}
	// swapPairs(n4).status()

	n3 := &ListNode{3, n4}
	swapPairs(n3).status()

	// n2 := &ListNode{2, n3}
	// n1 := &ListNode{1, n2}
	// swapPairs(n1).status()

}

func (n *ListNode) status() {
	fmt.Print("current list: ")
	tmp := n
	for tmp != nil {
		fmt.Printf("%v, ", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := head.Next
	prev := head
	current := head
	for current != nil && current.Next != nil {
		if prev != current {
			prev.Next = current.Next
		}
		tmp := current.Next.Next
		current.Next.Next = current
		current.Next = tmp

		prev = current
		current = current.Next
	}

	return res
}
