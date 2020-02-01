type MyQueue struct {
    inbox []int
    outbox []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    
    return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    
    this.inbox = append(this.inbox, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    
    if len(this.outbox) < 1 {
        for len(this.inbox) > 0 {
            item := this.inbox[len(this.inbox)-1]
            this.inbox = this.inbox[:len(this.inbox)-1]
            this.outbox = append(this.outbox, item)
        }
    }
    
    item := this.outbox[len(this.outbox) - 1]
    this.outbox = this.outbox[:len(this.outbox)-1]
    
    return item
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    
    if len(this.outbox) < 1 {
        for len(this.inbox) > 0 {
            item := this.inbox[len(this.inbox)-1]
            this.inbox = this.inbox[:len(this.inbox)-1]
            this.outbox = append(this.outbox, item)
        }
    }
    
    return this.outbox[len(this.outbox) - 1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    
    return len(this.inbox) == 0 && len(this.outbox) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
