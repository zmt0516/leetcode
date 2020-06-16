func myAtoi(str string) int {
	var r, i, sign int
	for i < len(str) {
		if str[i] == ' ' {
			i++
			continue
		}
		if str[i] == '-' || str[i] == '+' {
			if str[i] == '-' {
				sign = -1
			} else {
				sign = 1
			}
			i++
			if i >= len(str) || int(str[i])-int('9') > 0 || int(str[i])-int('0') < 0 {
				return 0
			} else {
				r += int(str[i]) - int('0')
			}
			for i++; i < len(str); i++ {
				if int(str[i])-int('9') <= 0 && int(str[i])-int('0') >= 0 {
					r *= 10
					r += int(str[i]) - int('0')
					if r > int(^uint32(0)>>1)+1 && sign == -1 {
						return -1 - int(^uint32(0)>>1)
					} else if r > int(^uint32(0)>>1) && sign == 1 {
						return int(^uint32(0) >> 1)
					}

				} else if sign == 1 {
					return r
				} else {
					return -r
				}

			}
			if sign == -1 {
				r = -r
			}
			return r

		} else if int(str[i])-int('9') > 0 || int(str[i])-int('0') < 0 {
			return 0
		} else {

			r += int(str[i]) - int('0')
			for i++; i < len(str); i++ {
				if int(str[i])-int('9') <= 0 && int(str[i])-int('0') >= 0 {
					r *= 10
					r += int(str[i]) - int('0')
					if r > int(^uint32(0)>>1) {
						return int(^uint32(0) >> 1)
					}
				} else {
					return r
				}
			}
			return r
		}
	}
	return r

}