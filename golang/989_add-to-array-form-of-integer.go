func addToArrayForm(A []int, K int) []int {

	n := 0
	for K > 0 {
		if n < len(A) {
			K += A[len(A)-1-n]
			A[len(A)-1-n] = K % 10
			K /= 10
		} else {
			A = append([]int{K % 10}, A...)
			K /= 10
		}
		n += 1

	}
	return A
}