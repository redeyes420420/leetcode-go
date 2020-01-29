type MinStack struct {
    backingArray []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    
    this.backingArray = append(this.backingArray, x)
}


func (this *MinStack) Pop()  {

    this.backingArray = this.backingArray[:len(this.backingArray)-1]
}


func (this *MinStack) Top() int {
    
    return this.backingArray[len(this.backingArray)-1]
}


func (this *MinStack) GetMin() int {

    min := math.MaxInt64
    for _, v := range this.backingArray {
        if v < min {
            min = v
        }
    }
    
    return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
