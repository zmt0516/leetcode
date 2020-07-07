func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
func Max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func simpleintervalIntersection(a []int, B [][]int, c *int) [][]int {
	var r [][]int
	for ; *c < len(B); *c++ {
		b := B[*c]
		switch {
		case a[1] < b[0]:
			return r
		case a[1] >= b[0] && b[1] >= a[0]:
			r = append(r, []int{Max(a[0], b[0]), Min(a[1], b[1])})
			//case b[1]<a[0]:
			//pass
			//return r

		}

	}
	return r

}

func intervalIntersection(A [][]int, B [][]int) [][]int {

	var r [][]int
	c := 0
	for _, v := range A {
		//fmt.Println(v,c)
		if c != 0 {
			c--
		}
		r = append(r, simpleintervalIntersection(v, B, &c)...)

	}
	return r

}