/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	inc := new(ListNode)
	head = inc

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			inc.Next = l2
			l2 = l2.Next
			inc = inc.Next
			if l2 == nil {
				inc.Next = l1
				return head.Next
			}
		} else {
			inc.Next = l1
			l1 = l1.Next
			inc = inc.Next
			if l1 == nil {
				inc.Next = l2
				return head.Next
			}
		}

	}
	if l1 == nil {
		return l2
	} else {
		return l1
	}
	return nil

}