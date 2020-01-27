/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    
    l3 := new(ListNode)
    cur := l3
    for l1 != nil && l2 != nil {
        cur.Next = new(ListNode)
        cur = cur.Next
        if l1.Val < l2.Val {
            cur.Val = l1.Val
            l1 = l1.Next
        } else {
            cur.Val = l2.Val
            l2 = l2.Next
        }
    }
    
    for l1 != nil {
        cur.Next = new(ListNode)
        cur = cur.Next
        cur.Val = l1.Val
        l1 = l1.Next
    }
    for l2 != nil {
        cur.Next = new(ListNode)
        cur = cur.Next
        cur.Val = l2.Val
        l2 = l2.Next
    }
    
    return l3.Next
}
