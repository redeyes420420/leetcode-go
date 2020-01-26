/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    var sl []*ListNode
    
    cur := head
    for cur != nil {
        sl = append(sl, cur)
        cur = cur.Next
    }
    
    theNode := sl[len(sl)-n]
    if theNode == head {
        head = head.Next
        return head
    }
    
    prev := sl[len(sl)-n-1]
    prev.Next = theNode.Next
    
    return head
}
