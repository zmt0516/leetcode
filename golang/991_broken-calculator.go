func brokenCalc(X int, Y int) int {
	n := 0
	for X < Y {
		X *= 2
		n++
	}
	X -= Y
	r := n
	for i := 0; i < n; i++ {
		r += X % 2
		X /= 2
	}
	r += X
	return r

}