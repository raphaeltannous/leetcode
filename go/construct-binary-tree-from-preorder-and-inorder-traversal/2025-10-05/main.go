import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
  if len(preorder) == 0 {
    return nil
  }

  root := &TreeNode{
    Val: preorder[0],
  }

  rootInorderIndex := slices.Index(inorder, root.Val)

  if rootInorderIndex != 0 {
    root.Left = buildTree(preorder[1:rootInorderIndex+1], inorder[:rootInorderIndex])
  } 
  root.Right = buildTree(preorder[rootInorderIndex+1:], inorder[rootInorderIndex+1:])

  return root
}
