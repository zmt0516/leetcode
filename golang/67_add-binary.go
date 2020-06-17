func addBinary(a string, b string) string {
	inc := 0
	n := 0
	r := ""
	for n < len(a) || n < len(b) {
		if n < len(a) && a[len(a)-1-n] == '1' {
			inc += 1
		}
		if n < len(b) && b[len(b)-1-n] == '1' {
			inc += 1
		}
		if inc%2 == 1 {
			r = "1" + r
		} else {
			r = "0" + r
		}
		inc /= 2
		n++

	}
	if inc == 1 {
		r = "1" + r
	}
	return r

}