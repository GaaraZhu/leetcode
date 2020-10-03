/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    current1 := l1
    current2 := l2
    var result *ListNode=nil
    var current *ListNode=result
    for (current1 != nil || current2 != nil) {
        if(current2==nil || (current1!=nil && current1.Val<current2.Val)) {
            if(current==nil) {
                result = &ListNode{Val: current1.Val, Next: nil}
                current = result
            } else {
                next := &ListNode{Val: current1.Val, Next: nil}
                current.Next = next
                current = next
            }
            current1 = current1.Next
            continue
        } else {
            if(current==nil) {
                result = &ListNode{Val: current2.Val, Next: nil}
                current = result
            } else {
                next := &ListNode{Val: current2.Val, Next: nil}
                current.Next = next
                current = next
            }
            current2 = current2.Next
        }
    }
    
    return result
}