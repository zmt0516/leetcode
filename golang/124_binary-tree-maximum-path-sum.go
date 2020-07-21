/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

var r int

func maxPathSum(root *TreeNode) int {
	r = MinInt
	getMax(root)
	return r

}

func getMax(root *TreeNode) int {

	if root == nil {
		return 0
	}
	//fmt.Println(r)
	x := getMax(root.Left)
	y := getMax(root.Right)
	if root.Val+x+y > r {
		r = root.Val + x + y
	}
	//fmt.Println(x,y,r)
	return Max(root.Val, Max(root.Val+x, root.Val+y))

}

func Max(x int, y int) int {

	if x < y {
		if y > r {
			r = y
		}
		return y
	}
	if x > r {
		r = x
	}
	return x

}