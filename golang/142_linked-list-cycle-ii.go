/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	inc1 := new(ListNode)
	inc2 := new(ListNode)
	inc := new(ListNode)
	inc1 = head
	if inc1 == nil {
		return nil
	}
	inc2 = head.Next

	var n, n2, k int

	for n = 0; inc1 != inc2; n++ {
		if inc2 == nil {
			return nil
		}
		inc2 = inc2.Next
		if n%2 == 1 {
			inc1 = inc1.Next

		}

	}

	//n2 = n / 2
	n += 1

	n2 = n/2 + 1
	inc = head
	inc1 = head
	inc2 = head

	//println(n)
	for i := 0; i < n; i++ {

		inc2 = inc2.Next
		if inc1 == inc2 {
			return head
		}
	}

	for {
		k = 0
		for i := 0; i < n2; i++ {
			inc1 = inc1.Next
			inc2 = inc2.Next
		}
		for i := 0; i < n; i++ {

			inc2 = inc2.Next
			if inc1 == inc2 {
				k = 1
			}
		}
		if k == 0 {
			inc = inc1
			inc1 = inc
			inc2 = inc
		} else {
			inc1 = inc
			inc2 = inc
			n2 = n2 / 2
			//println(n2)

		}
		if n2 == 0 {
			inc1 = inc1.Next
			inc2 = inc2.Next

			k = 0
			for i := 0; i < n; i++ {
				inc2 = inc2.Next
				if inc1 == inc2 {
					k = 1
				}

			}
			if k == 1 {
				return inc.Next
			} else {
				n2 = 1
			}

		}

	}
	return nil

}