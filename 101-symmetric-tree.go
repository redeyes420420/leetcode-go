/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    
    if root == nil {
        return true
    }
    
    return helper(root.Left, root.Right)
}

func helper(node1, node2 *TreeNode) bool {
    
    if node1 == nil && node2 == nil {
        return true
    } else if node1 == nil || node2 == nil{
        return false
    }
    
    return node1.Val == node2.Val && helper(node1.Left, node2.Right) && helper(node2.Left, node1.Right)
}
