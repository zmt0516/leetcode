/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	rl := []*TreeNode{root}
	for i := 0; i < len(rl); i++ {
		if rl[i].Left != nil {
			rl = append(rl, rl[i].Left)
		}
		if rl[i].Right != nil {
			rl = append(rl, rl[i].Right)
		}
		for j := 0; j < i; j++ {
			if rl[i].Val+rl[j].Val == k {
				return true
			}

		}

	}
	return false

}