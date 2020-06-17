/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var inc int
	var l3 = new(ListNode)
	l4 := l3
	for l1 != nil || l2 != nil || inc != 0 {
		if l1 != nil {
			inc += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			inc += l2.Val
			l2 = l2.Next
		}
		l4.Next = &ListNode{Val: inc % 10}
		l4 = l4.Next
		inc /= 10

	}
	return l3.Next

}