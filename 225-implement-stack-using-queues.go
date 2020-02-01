type MyStack struct {
    queue []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    
    return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    
    this.queue = append(this.queue, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    
    var bufQ []int
    for len(this.queue) > 1 {
        item := this.queue[0]
        this.queue = this.queue[1:]
        bufQ = append(bufQ, item)
    }
    item := this.queue[0]
    this.queue = bufQ
    return item
}


/** Get the top element. */
func (this *MyStack) Top() int {
    
    var bufQ []int
    for len(this.queue) > 1 {
        item := this.queue[0]
        this.queue = this.queue[1:]
        bufQ = append(bufQ, item)
    }
    item := this.queue[0]
    bufQ = append(bufQ, item)
    this.queue = bufQ
    return item
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    
    return len(this.queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
