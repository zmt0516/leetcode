/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, L int, R int) *TreeNode {

	if root == nil {
		return root
	}

	if root.Val >= L && root.Val <= R {

		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root

	}

	temp := trimBST(root.Left, L, R)
	if temp == nil {
		temp = trimBST(root.Right, L, R)
	}
	return temp

}