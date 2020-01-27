/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
   
    if p == nil && q == nil {
        return true
    } else if p == nil || q == nil || q.Val != p.Val {
        return false
    } else {
        return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
    }
}
