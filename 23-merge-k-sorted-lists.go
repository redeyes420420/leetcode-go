/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import (
    "container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *MinHeap) Push(v interface{}) {
    
    *h = append(*h, v.(int))
}
func (h *MinHeap) Pop() interface{} {
    
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    
    return x
}

func mergeKLists(lists []*ListNode) *ListNode {
    
    result := &ListNode{0, nil}
    h := &MinHeap{}
    heap.Init(h)
    
    for _, head := range lists {
        cur := head
        for cur != nil {
            heap.Push(h, cur.Val)
            cur = cur.Next
        }
    }
    if h.Len() == 0 {
        return nil
    }
    
    cur := result
    for h.Len() > 0 {
        cur.Val = heap.Pop(h).(int)
        if h.Len() == 0 {
            break
        }
        cur.Next = &ListNode{0, nil}
        cur = cur.Next
    }
    
    return result
}
