func numDistinct(s string, t string) int {
	var r int
	lens := len(s)
	lent := len(t)
	if lens < lent {
		return 0
	}
	if lent == 0 {
		return 1
	}
	for i := 0; i <= lens-lent; i++ {
		if s[i] == t[0] {
			r += numDistinct(s[i+1:], t[1:])
		}

	}

	return r
}