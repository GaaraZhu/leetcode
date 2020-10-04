/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    current := head
    for current != nil && current.Next != nil {
        var tmp *ListNode
        for tmp = current.Next;tmp != nil && tmp.Val == current.Val;{
            tmp = tmp.Next
        }
        current.Next = tmp
        current = current.Next
    }
    
    return head
}