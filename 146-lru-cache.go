type Node struct {
    key int
    val int
    prev *Node
    next *Node
}

func newNode(key, val int) *Node {
    
    n := new(Node)
    n.key = key
    n.val = val
    
    return n
}

type LRUCache struct {
    
    keys map[int]*Node
    capacity int
    size int
    head *Node
    tail *Node
}


func Constructor(capacity int) LRUCache {
    
    lru := LRUCache{
        keys: make(map[int]*Node),
        capacity: capacity,
        size: 0,
        head: newNode(-1, -1),
        tail: newNode(-1, -1),
    }
    lru.head.next = lru.tail
    lru.tail.prev = lru.head
    
    return lru
}

func (*LRUCache) removeNode(n *Node) {
    
    n.prev.next = n.next
    n.next.prev = n.prev
}

func (this *LRUCache) addAtTail(n *Node) {
    
    realTail := this.tail.prev
    
    realTail.next = n
    n.prev = realTail
    
    n.next = this.tail
    this.tail.prev = n
    
}

func (this *LRUCache) Get(key int) int {
    
    n, f := this.keys[key]
    if !f {
        return -1
    }
    
    this.removeNode(n)
    this.addAtTail(n)
    
    return n.val
}


func (this *LRUCache) Put(key int, value int)  {
    
    n, f := this.keys[key]
    if !f {
        n = newNode(key, value)
        this.keys[key] = n
        this.addAtTail(n)
        this.size++
    } else {
        this.removeNode(n)
        n.val = value
        this.addAtTail(n)
    }
    
    if this.size > this.capacity {
        realHead := this.head.next
        delete(this.keys, realHead.key)
        this.removeNode(realHead)
        this.size--
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
