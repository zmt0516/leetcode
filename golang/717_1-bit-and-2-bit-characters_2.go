func isOneBitCharacter(bits []int) bool {
	res := 0
	for res < len(bits)-1 {
		if bits[res] == 1 {
			res += 2
		} else {
			res++
		}
	}
	if res == len(bits) {
		return false
	}
	return true

}