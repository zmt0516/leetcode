/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(inorder) != 0 {
		root := new(TreeNode)
		root.Val = preorder[0]
		var i, v int
		for i, v = range inorder {
			if v == root.Val {
				break
			}

		}
		root.Left = buildTree(preorder[1:i+1], inorder[:i])
		root.Right = buildTree(preorder[i+1:], inorder[i+1:])
		return root

	}
	return nil

}