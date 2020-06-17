func plusOne(digits []int) []int {

	inc := 1
	for i := range digits {
		inc += digits[len(digits)-1-i]
		digits[len(digits)-1-i] = inc % 10
		inc = inc / 10
	}
	if inc != 0 {
		digits = append([]int{1}, digits...)
	}
	return digits

}