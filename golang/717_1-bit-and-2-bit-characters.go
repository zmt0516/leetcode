func isOneBitCharacter(bits []int) bool {
	res := false
	for i := 0; i < len(bits); i++ {
		if bits[i] == 0 {
			res = true
		} else {
			res = false
			i++

		}
	}
	return res

}