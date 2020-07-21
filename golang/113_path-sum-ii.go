import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, sum int) [][]int {
	//r:=make([][]int,0)
	var r [][]int
	var trial []int
	//trial:=make([]int,0)
	Sum(root, sum, trial, &r)
	return r

}

func Sum(root *TreeNode, sum int, trial []int, r *[][]int) {

	if root == nil {
		return
	}

	trial = append(trial, root.Val)
	if root.Right == nil && root.Left == nil && root.Val == sum {

		newtrial := make([]int, len(trial))
		copy(newtrial, trial)
		fmt.Println(newtrial)
		*r = append(*r, newtrial)
	}

	Sum(root.Left, sum-root.Val, trial, r)

	Sum(root.Right, sum-root.Val, trial, r)

}