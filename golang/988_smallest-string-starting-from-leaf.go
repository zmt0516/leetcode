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

var r []int

//var n int
func smallestFromLeaf(root *TreeNode) string {
	r = make([]int, 0)
	//n=MaxInt
	fromLeaf(root, r)

	s := ""
	for _, v := range r {

		s += string('a' + v)

	}
	return s

}

func fromLeaf(root *TreeNode, arr []int) {
	if root == nil {

		return

	}
	var arr2 []int
	arr2 = append(arr2, root.Val)
	arr2 = append(arr2, arr...)

	if root.Left == nil && root.Right == nil {
		leaf(arr2)

	}
	fromLeaf(root.Left, arr2)
	fromLeaf(root.Right, arr2)

}

func leaf(arr2 []int) {
	if len(r) == 0 {
		r = make([]int, 0)
		r = append(r, arr2...)
		return

	}

	for i := 0; i < len(arr2) && i < len(r); i++ {
		if arr2[i] < r[i] {
			r = make([]int, 0)
			r = append(r, arr2...)
			return
		}
		if arr2[i] > r[i] {
			return
		}

	}
	if len(arr2) < len(r) {
		r = make([]int, 0)
		r = append(r, arr2...)
	}

}
 
 