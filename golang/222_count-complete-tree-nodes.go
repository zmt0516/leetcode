/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	var r int
	nodes(root, &r)
	return r

}

func nodes(root *TreeNode, r *int) {
	if root == nil {
		return
	}
	*r++
	nodes(root.Left, r)
	nodes(root.Right, r)

}