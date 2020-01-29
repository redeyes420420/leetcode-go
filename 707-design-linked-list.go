type Node struct {
    val int
    next *Node
}

func newNode(val int) *Node {
    
    n := new(Node)
    n.val = val
    return n
}

type MyLinkedList struct {
    head *Node
    length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {

    return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    
    if index < 0 || index >= this.length {
        return -1
    }
    
    cur := this.head
    for index > 0 {
        cur = cur.next
        index--
    }
    if cur == nil {
        return -1
    }
    return cur.val
    
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    
    n := newNode(val)
    
    if this.head == nil {
        this.head = n
        this.length++
        return
    }
    
    n.next = this.head
    this.head = n
    this.length++

}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    
    n := newNode(val)
    
    if this.head == nil {
        this.head = n
        this.length++
        return
    }
    
    cur := this.head
    for cur.next != nil {
        cur = cur.next
    }
    cur.next = n
    this.length++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    
    n := newNode(val)
    
    if index < 0 || index > this.length {
        return
    }
    
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    
    if index == this.length {
        this.AddAtTail(val)
        return
    }
    
    cur := this.head
    for index > 1 {
        cur = cur.next
        index--
    }
    
    n.next = cur.next
    cur.next = n
    this.length++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    
    if index < 0 || index >= this.length {
        return
    }
    
    if index == 0 {
        this.head = this.head.next
        this.length--
        return
    }
    
    cur := this.head
    for index > 1 {
        cur = cur.next
        index--
    }
    
    if cur.next == nil {
        return
    } 
    cur.next = cur.next.next 
}

func (this *MyLinkedList) toStr() string {
    
    var str string
    cur := this.head
    for cur != nil {
        if cur.next != nil {
            str += fmt.Sprintf("%d -> ", cur.val)
        } else {
            str += fmt.Sprintf("%d", cur.val)
        }
        cur = cur.next
    }
    
    return str
}
/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
