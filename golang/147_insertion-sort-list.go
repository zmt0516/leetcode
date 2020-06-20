/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	inc := head
	//inc2 := head
	for inc != nil {
		inc2 := head

		for inc2 != inc && inc.Val > inc2.Val {
			inc2 = inc2.Next
		}

		incVal := 0
		inc2.Val, incVal = inc.Val, inc2.Val
		inc2 = inc2.Next
		for inc2 != inc.Next {
			inc2.Val, incVal = incVal, inc2.Val
			inc2 = inc2.Next
		}

		inc = inc.Next
	}
	return head

}