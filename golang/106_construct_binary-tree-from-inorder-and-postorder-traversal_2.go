/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	var i, v int
	if len(inorder) != 0 {
		root := new(TreeNode)
		root.Val = postorder[len(postorder)-1]

		for i, v = range inorder {
			if v == root.Val {
				break
			}
		}
		root.Left = buildTree(inorder[:i], postorder[:i])
		root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
		return root

	}
	return nil

}