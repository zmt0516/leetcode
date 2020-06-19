/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	head := new(ListNode)
	inc := new(ListNode)
	head = inc
	for i := 0; i < len(lists); {
		if lists[i] == nil {
			lists = append(lists[:i], lists[i+1:]...)
		} else {
			i++
		}
	}
	for len(lists) > 0 {
		vmin := int(^uint32(0) >> 1)
		v := new(ListNode)
		vi := 0
		i := 0
		for i, v = range lists {
			if v.Val < vmin {
				inc.Next = v
				vi = i
				vmin = v.Val
			}
		}
		if inc.Next.Next == nil {
			lists = append(lists[:vi], lists[vi+1:]...)
		} else {
			lists[vi] = lists[vi].Next
		}
		inc = inc.Next

	}
	return head.Next

}