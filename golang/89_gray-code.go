func grayCode(n int) []int {
	var r []int
	n2 := 1
	r = append(r, 0)

	for i := 0; i < n; i++ {
		for j := n2 - 1; j >= 0; j-- {
			r = append(r, r[j]+n2)
		}
		n2 *= 2
	}
	return r

}