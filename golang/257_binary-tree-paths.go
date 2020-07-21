/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	var r []string
	if root == nil {
		return r
	}
	s := strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		r = append(r, s)

	}

	bPath(root.Left, s, &r)
	bPath(root.Right, s, &r)

	return r
}

func bPath(root *TreeNode, s string, r *[]string) {

	if root == nil {
		return
	}
	s += "->" + strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		*r = append(*r, s)

	}

	bPath(root.Left, s, r)
	bPath(root.Right, s, r)

}