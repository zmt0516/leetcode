func canThreePartsEqualSum(A []int) bool {

	if len(A) < 3 {
		return false
	}
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	onepart := sum / 3
	sum = 0
	n := 0
	for _, v := range A {
		sum += v
		if sum == onepart {
			n++
			if n > 2 {
				return true
			}
			sum -= onepart
		}
	}
	return false

}