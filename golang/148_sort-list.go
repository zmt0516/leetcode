/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortListQ(a *ListNode, left *ListNode, right *ListNode) {
	if left == right || left == right.Next {
		return
	}

	inc := new(ListNode)
	inc.Next = left
	for i := left.Next; ; i = i.Next {
		if left.Val >= i.Val {
			inc = inc.Next
			i.Val, inc.Next.Val = inc.Next.Val, i.Val
		}
		if i == right {
			break
		}
	}
	left.Val, inc.Next.Val = inc.Next.Val, left.Val
	sortListQ(a, left, inc)
	sortListQ(a, inc.Next.Next, right)

}

func sortList(head *ListNode) *ListNode {
	left := head
	right := head
	for i := head; i != nil; i = i.Next {
		if i.Next == nil {
			right = i
			break
		}
	}
	sortListQ(head, left, right)
	return head

}