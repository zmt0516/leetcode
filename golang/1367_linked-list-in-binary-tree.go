/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}

	if head.Val == root.Val {
		return SubPath(head.Next, root.Right) || SubPath(head.Next, root.Left) || isSubPath(head, root.Right) || isSubPath(head, root.Left)
	}
	return isSubPath(head, root.Right) || isSubPath(head, root.Left)

}
func SubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val == root.Val {
		return SubPath(head.Next, root.Left) || SubPath(head.Next, root.Right)
	}
	return false

}