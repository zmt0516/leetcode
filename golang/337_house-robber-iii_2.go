/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {

	if root == nil {
		return 0
	}

	r0, r1 := rob3(root)
	return Max(r0, r1)

}

func rob3(root *TreeNode) (int, int) {

	if root == nil {
		return 0, 0
	}

	l0, l1 := rob3(root.Left)
	r0, r1 := rob3(root.Right)

	return root.Val + l1 + r1, Max(l0, l1) + Max(r0, r1)

}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}