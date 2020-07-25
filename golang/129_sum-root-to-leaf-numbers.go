/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var r int

func sumNumbers(root *TreeNode) int {
	r = 0
	Numbers(root, 0)
	return r

}

func Numbers(root *TreeNode, x int) {
	if root == nil {
		return
	}

	x *= 10
	x += root.Val

	if root.Left == nil && root.Right == nil {
		r += x
		return

	}

	Numbers(root.Left, x)
	Numbers(root.Right, x)

}