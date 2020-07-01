import "strconv"

func restoreIpAddresses(s string) []string {
	var r []string
	if s == "" {
		return r
	}
	for i := 1; i < 4; i++ {
		if s[0] == '0' && i > 1 {
			break
		}
		rs, _ := strconv.Atoi(s[:i])
		if rs > 255 || len(s)-i < 3 {
			break
		}
		for j := 1; j < 4; j++ {
			if s[i] == '0' && j > 1 {
				break
			}
			rs, _ = strconv.Atoi(s[i : i+j])
			//fmt.Println(rs)
			if rs > 255 || len(s)-i-j < 2 {
				break
			}

			for k := 1; k < 4; k++ {
				if s[i+j] == '0' && k > 1 {
					break
				}
				rs, _ = strconv.Atoi(s[i+j : i+j+k])

				if rs > 255 || len(s)-i-j-k < 1 {
					break
				}
				if s[i+j+k] == '0' && len(s)-i-j-k > 1 {
					continue
				}
				rs, _ = strconv.Atoi(s[i+j+k:])
				if rs > 255 {
					continue
				}

				r = append(r, s[:i]+"."+s[i:i+j]+"."+s[i+j:i+j+k]+"."+s[i+j+k:])

			}

		}
	}
	return r

}