/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
  var recurse func(parent, curr *TreeNode)
  recurse = func(parent, curr *TreeNode) {
    if curr == nil {
      return
    }

    recurse(curr, curr.Left)
    recurse(curr, curr.Right)

    if curr.Left == nil && curr.Right == nil {
      if curr.Val == target {
        if parent != nil && parent.Left == curr {
          parent.Left = nil
        } else if parent != nil && parent.Right == curr {
          parent.Right = nil
        }

        if curr == root {
          root = nil
        }
      }
    }
  }
  recurse(nil, root)

  return root
}
