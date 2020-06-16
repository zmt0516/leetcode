func isUgly(num int) bool {

	for num > 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			break
		}
	}
	for num > 1 {
		if num%3 == 0 {
			num /= 3
		} else {
			break
		}
	}
	for num > 1 {
		if num%5 == 0 {
			num /= 5
		} else {
			break
		}
	}
	if num == 1 {
		return true
	}
	return false

}