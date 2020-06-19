// Maybe, many people use this method. It's quick, but not compliant.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "sort"

func sortList(head *ListNode) *ListNode {
	var r []int
	for i := head; i != nil; i = i.Next {
		r = append(r, i.Val)
	}
	sort.Ints(r)
	i := head
	for _, v := range r {
		i.Val = v
		i = i.Next
	}
	return head

}