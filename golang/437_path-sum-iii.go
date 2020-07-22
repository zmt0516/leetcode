/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var r int

func pathSum(root *TreeNode, sum int) int {
	r = 0
	if root == nil {
		return r
	}
	var ps []int
	path(root, sum, ps)

	return r

}

func path(root *TreeNode, sum int, ps []int) {
	var ps2 []int
	if root == nil {
		return
	}

	for i := range ps {
		ps2 = append(ps2, ps[i]+root.Val)
		if ps2[i] == sum {
			r++
		}
	}
	ps2 = append(ps2, root.Val)
	if root.Val == sum {
		r++
	}
	path(root.Left, sum, ps2)
	path(root.Right, sum, ps2)

}