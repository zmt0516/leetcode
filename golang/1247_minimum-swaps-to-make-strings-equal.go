func minimumSwap(s1 string, s2 string) int {
	m := 0
	n := 0
	for i, v := range s1 {
		switch v {
		case 'x':
			if s2[i] == 'y' {
				m++
			}
		case 'y':
			if s2[i] == 'x' {
				n++
			}

		}

	}
	r := 0
	r += m/2 + n/2
	m %= 2
	n %= 2
	switch m + n {
	case 1:
		return -1
	case 2:
		return r + 2
	}
	return r

}