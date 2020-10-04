/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head==nil || head.Next==nil || head.Next.Next==nil {
        return false
    }
    pointerOne := head.Next
    pointerTwo := head.Next.Next
    
    for(pointerOne!=nil && pointerTwo!=nil) {
        if pointerOne==pointerTwo {
            return true
        }
        if (pointerOne.Next==nil || pointerTwo.Next==nil || pointerTwo.Next.Next==nil) {
            return false
        }
        
        pointerOne = pointerOne.Next
        pointerTwo = pointerTwo.Next.Next
    }
    
    return false
}