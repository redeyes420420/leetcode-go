/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    sorted []int
}


func Constructor(root *TreeNode) BSTIterator {
    
    var iter BSTIterator
    iter.inOrder(root)
    fmt.Println(iter.sorted)
    
    return iter
}

func (this *BSTIterator) inOrder(cur *TreeNode) {
    
    if cur == nil {
        return
    }
    
    this.inOrder(cur.Left)
    this.sorted = append(this.sorted, cur.Val)
    this.inOrder(cur.Right)
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    
    n := this.sorted[0]
    this.sorted = this.sorted[1:]
    return n
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    
    return len(this.sorted) != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
