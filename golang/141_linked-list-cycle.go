/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	for head != nil {

		if head.Next != head {
			head, head.Next = head.Next, head

		} else {
			return true
		}
	}
	return false
}