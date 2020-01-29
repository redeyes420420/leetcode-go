/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    
    if head == nil || head.Next == nil {
        return head
    }
    
    a, b := head, head.Next
    b.Next, a.Next = a, swapPairs(b.Next)
    
    return b
}
