/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var leftNode *ListNode
	var rightNode *ListNode
	var currentLeftNode *ListNode
	var currentRightNode *ListNode
	for head != nil {
		if head.Val < x {
			if leftNode == nil {
				leftNode = &ListNode{
					Val:  head.Val,
					Next: nil,
				}
				currentLeftNode = leftNode
			} else {
				currentLeftNode.Next = &ListNode{
					Val:  head.Val,
					Next: nil,
				}
				currentLeftNode = currentLeftNode.Next
			}
		} else {
			if rightNode == nil {
				rightNode = &ListNode{
					Val:  head.Val,
					Next: nil,
				}
				currentRightNode = rightNode
			} else {
				currentRightNode.Next = &ListNode{
					Val:  head.Val,
					Next: nil,
				}
				currentRightNode = currentRightNode.Next
			}
		}
		head = head.Next
	}

	if leftNode != nil {
		currentLeftNode.Next = rightNode
		return leftNode
	} else {
		return rightNode
	}
}