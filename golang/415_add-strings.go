func addStrings(num1 string, num2 string) string {
	inc := 0
	r := ""
	for n := 0; n < len(num1) || n < len(num2); n++ {
		if n < len(num1) {
			inc += int(num1[len(num1)-1-n]) - int('0')
		}
		if n < len(num2) {
			inc += int(num2[len(num2)-1-n]) - int('0')
		}
		r = string(inc%10+int('0')) + r
		inc /= 10
	}
	if inc != 0 {
		r = "1" + r
	}
	return r
}