func add(num1 string, num2 string) string {
	r := num1
	inc := 0
	for i := 1; i <= len(r) || i <= len(num2); i++ {
		if i <= len(r) {
			inc += int(r[len(r)-i]) - int('0')
		}
		if i <= len(num2) {
			inc += int(num2[len(num2)-i]) - int('0')
		}
		if i <= len(r) {
			r = r[:len(r)-i] + string(int('0')+inc%10) + r[len(r)-i+1:]
		} else {
			r = string(int('0')+inc%10) + r
		}
		inc /= 10

	}
	if inc == 1 {
		r = "1" + r
	}
	return r
}

func multiplyn(num1 string, num2 int) string {
	inc := 0
	r := num1
	for i := 1; i <= len(r); i++ {
		inc += (int(r[len(r)-i]) - int('0')) * num2
		r = r[:len(r)-i] + string(int('0')+inc%10) + r[len(r)-i+1:]
		inc /= 10
	}
	if inc > 0 {
		r = string(int('0')+inc) + r
	}
	return r
}

func multiply(num1 string, num2 string) string {
	if num2 == "0" || num1 == "0" {
		return "0"
	}
	r := "0"
	zero := ""
	for i := range num2 {
		r = add(r, multiplyn(num1+zero, int(num2[len(num2)-1-i])-int('0')))
		zero += "0"
	}

	return r
}