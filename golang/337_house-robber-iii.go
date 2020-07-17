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

	return Max(root.Val+rob2(root.Left)+rob2(root.Right), rob2(root))

}

func rob2(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return rob(root.Left) + rob(root.Right)

}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}