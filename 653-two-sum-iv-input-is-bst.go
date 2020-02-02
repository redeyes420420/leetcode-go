/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    
    var f bool
    m := make(map[int]bool, 0)
    
    cb := func(v int) {
        needed := k - v
        if _, found := m[needed]; found {
            f = true
            return
        
        } else {
            m[v] = true
        } 
    }
    
    preOrder(root, cb)
    
    return f
}

func preOrder(node *TreeNode, cb func(v int)) {
    
    if node == nil {
        return
    }
    
    cb(node.Val)
    preOrder(node.Left, cb)
    preOrder(node.Right, cb)
}
